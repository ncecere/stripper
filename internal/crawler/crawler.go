package crawler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"stripper/internal/ai"
	"stripper/internal/database"
	"stripper/internal/storage"
	"stripper/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gocolly/colly/v2"
)

// Crawler handles the web crawling functionality
type Crawler struct {
	client         *http.Client
	baseURL        *url.URL
	depth          int
	format         string
	force          bool
	ignore         []string
	outputDir      string
	storage        storage.Storage
	db             *database.DB
	ui             *tea.Program
	rescanInterval time.Duration
	readerAPIURL   string
	parallelism    int
	aiEnabled      bool
	aiClient       *ai.Client
	systemPrompt   string
}

// Options configures the crawler behavior
type Options struct {
	URL            string
	Depth          int
	Format         string
	Force          bool
	Ignore         []string
	OutputDir      string
	RescanInterval time.Duration
	ReaderAPIURL   string
	Parallelism    int
	AI             struct {
		Enabled      bool
		Endpoint     string
		APIKey       string
		Model        string
		SystemPrompt string
	}
}

// New creates a new Crawler instance
func New(opts Options) (*Crawler, error) {
	baseURL, err := url.Parse(opts.URL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	// Initialize storage
	store, err := storage.NewFileStorage(opts.OutputDir)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize storage: %w", err)
	}

	// Initialize database
	dbPath := path.Join(opts.OutputDir, "crawler.db")
	db, err := database.New(dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	// Set default Reader API URL if not provided
	readerAPIURL := opts.ReaderAPIURL
	if readerAPIURL == "" {
		readerAPIURL = "https://read.tabnot.space"
	}

	// Create crawler instance
	c := &Crawler{
		client:         &http.Client{},
		baseURL:        baseURL,
		depth:          opts.Depth,
		format:         opts.Format,
		force:          opts.Force,
		ignore:         opts.Ignore,
		outputDir:      opts.OutputDir,
		storage:        store,
		db:             db,
		rescanInterval: opts.RescanInterval,
		readerAPIURL:   readerAPIURL,
		parallelism:    opts.Parallelism,
		aiEnabled:      opts.AI.Enabled,
		systemPrompt:   opts.AI.SystemPrompt,
	}

	// Initialize AI client if enabled
	if c.aiEnabled {
		if opts.AI.APIKey == "" {
			return nil, fmt.Errorf("AI API key is required when AI is enabled")
		}
		c.aiClient = ai.New(ai.Options{
			Endpoint: opts.AI.Endpoint,
			APIKey:   opts.AI.APIKey,
			Model:    opts.AI.Model,
		})
	}

	// Initialize TUI
	c.ui = tui.New(db)

	return c, nil
}

// Start begins the crawling process
func (c *Crawler) Start() error {
	var wg sync.WaitGroup
	errChan := make(chan error, 1)
	doneChan := make(chan bool, 1)

	// Start UI in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := c.ui.Start(); err != nil {
			fmt.Printf("Error starting UI: %v\n", err)
		}
	}()

	// Start crawling in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(doneChan)

		// First phase: Collect links using colly
		if err := c.collectLinks(); err != nil {
			errChan <- fmt.Errorf("error collecting links: %w", err)
			return
		}

		// Second phase: Process collected links using Reader API
		if err := c.processLinks(); err != nil {
			errChan <- fmt.Errorf("error processing links: %w", err)
			return
		}

		doneChan <- true
	}()

	// Start ticker for UI updates
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	// Wait for completion or error
	for {
		select {
		case t := <-ticker.C:
			c.ui.Send(tui.TickMsg(t))
		case err := <-errChan:
			c.ui.Quit()
			wg.Wait()
			return err
		case <-doneChan:
			c.ui.Quit()
			wg.Wait()
			return nil
		}
	}
}

// collectLinks uses colly to find all links on the site
func (c *Crawler) collectLinks() error {
	// Create collector without depth limit since we'll handle it ourselves
	collector := colly.NewCollector(
		colly.AllowedDomains(c.baseURL.Host),
		colly.Async(true),
	)

	// Add rate limiting
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: c.parallelism,
		Delay:       1 * time.Second,
	})

	debugf("Starting link collection for %s with depth %d", c.baseURL.String(), c.depth)

	// Queue the initial URL at depth 0
	if err := c.db.QueueLink(c.baseURL.String(), 0); err != nil {
		return fmt.Errorf("failed to queue initial URL: %w", err)
	}

	debugf("Initial URL queued: %s", c.baseURL.String())

	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link == "" {
			return
		}

		// Parse the link
		parsedLink, err := url.Parse(link)
		if err != nil {
			debugf("Error parsing URL %s: %v", link, err)
			return
		}

		// Only process links from the same domain as the base URL
		if parsedLink.Host != c.baseURL.Host {
			debugf("Skipping external domain: %s", link)
			return
		}

		// Skip ignored extensions
		if shouldIgnoreURL(link, c.ignore) {
			debugf("Ignoring URL: %s", link)
			return
		}

		// Calculate depth based on the parent request
		depth := e.Request.Depth + 1

		// Only queue links if we haven't reached max depth
		if depth <= c.depth {
			if err := c.db.QueueLink(link, depth); err != nil {
				debugf("Error queueing link %s: %v", link, err)
			} else {
				debugf("Queued link: %s (depth: %d)", link, depth)
				// Visit this URL to find more links, but only if we haven't reached max depth
				if depth < c.depth {
					e.Request.Visit(link)
				}
			}
		} else {
			debugf("Skipping link due to depth limit: %s (depth: %d)", link, depth)
		}
	})

	collector.OnError(func(r *colly.Response, err error) {
		debugf("Error collecting from %s: %v", r.Request.URL, err)
	})

	err := collector.Visit(c.baseURL.String())
	if err != nil {
		return err
	}

	// Wait for all collectors to finish
	collector.Wait()
	return nil
}

// processLinks processes queued links using the Reader API
func (c *Crawler) processLinks() error {
	const (
		batchSize      = 5 // Reduced batch size
		delay          = 1 * time.Second
		maxRetries     = 5                // Increased retries
		aiRateLimit    = 5 * time.Second  // Increased delay between AI requests
		backoffInitial = 5 * time.Second  // Increased initial backoff
		backoffMax     = 60 * time.Second // Increased max backoff
	)

	// Create a rate limiter for AI requests
	aiLimiter := time.NewTicker(aiRateLimit)
	defer aiLimiter.Stop()

	for {
		// Get next batch of links
		links, err := c.db.GetNextBatch(batchSize)
		if err != nil {
			return fmt.Errorf("error getting next batch: %w", err)
		}

		if len(links) == 0 {
			break
		}

		debugf("Starting batch processing with size %d", batchSize)

		// Create a channel to limit concurrent goroutines
		sem := make(chan bool, c.parallelism)
		var wg sync.WaitGroup
		errChan := make(chan error, len(links))

		// Process links in parallel
		for _, link := range links {
			wg.Add(1)
			sem <- true // Acquire semaphore
			go func(link database.Link) {
				defer wg.Done()
				defer func() { <-sem }() // Release semaphore

				debugf("Processing link: %s (depth: %d)", link.URL, link.Depth)

				// Always collect links from pages we visit to find new content
				if err := c.collectLinksFromURL(link.URL, link.Depth); err != nil {
					debugf("Error collecting links from %s: %v", link.URL, err)
				}

				// Check if we should recrawl content
				shouldCrawl, err := c.db.ShouldRecrawl(link.URL, c.force, c.rescanInterval)
				if err != nil {
					c.db.UpdateLinkStatus(link.URL, "failed", err)
					return
				}

				if !shouldCrawl {
					debugf("Skipping recent URL: %s (last crawled: %s)", link.URL, link.LastCrawled)
					return
				}

				debugf("Fetching content for URL: %s", link.URL)

				// Fetch content using Reader API
				content, err := c.fetch(link.URL)
				if err != nil {
					c.db.UpdateLinkStatus(link.URL, "failed", err)
					errChan <- err
					return
				}

				// Generate AI summary first if enabled
				var aiSummary string
				if c.aiEnabled && c.aiClient != nil {
					debugf("Attempting AI summary for %s", link.URL)

					// Wait for rate limiter
					<-aiLimiter.C

					// Try with exponential backoff
					backoff := backoffInitial
					for retries := 0; retries < maxRetries; retries++ {
						summary, err := c.aiClient.Summarize(content, c.systemPrompt)
						if err != nil {
							if strings.Contains(err.Error(), "429") {
								debugf("Rate limited, waiting %v before retry %d for %s", backoff, retries+1, link.URL)
								time.Sleep(backoff)
								backoff *= 2
								if backoff > backoffMax {
									backoff = backoffMax
								}
								continue
							}
							debugf("Error generating AI summary for %s: %v", link.URL, err)
							break
						}
						debugf("Successfully generated AI summary for %s (%d chars)", link.URL, len(summary))
						aiSummary = summary
						break
					}
				} else {
					debugf("Skipping AI summary for %s (enabled: %v, client: %v)", link.URL, c.aiEnabled, c.aiClient != nil)
				}

				// If AI is enabled but we failed to get a summary, mark as failed
				if c.aiEnabled && c.aiClient != nil && aiSummary == "" {
					c.db.UpdateLinkStatus(link.URL, "failed", fmt.Errorf("failed to generate AI summary"))
					return
				}

				// Store original content
				if err := c.storage.Save(link.URL, content, c.format); err != nil {
					c.db.UpdateLinkStatus(link.URL, "failed", err)
					errChan <- err
					return
				}

				// Store AI summary if generated
				if aiSummary != "" {
					// Create AI output directory
					aiOutputDir := path.Join(c.outputDir, "ai")
					if err := os.MkdirAll(aiOutputDir, 0755); err != nil {
						debugf("Error creating AI output directory: %v", err)
						return
					}

					// Create flat file name for AI output
					fileName := strings.TrimPrefix(link.URL, c.baseURL.String())
					fileName = strings.ReplaceAll(fileName, "/", "_")
					if fileName == "" {
						fileName = "index"
					}
					fileName = strings.TrimSuffix(fileName, "_") + ".md"

					// Save directly to ai directory
					debugf("Saving AI summary to: ai/%s (from URL: %s)", fileName, link.URL)
					if err := os.WriteFile(path.Join(aiOutputDir, fileName), []byte(aiSummary), 0644); err != nil {
						debugf("Error saving AI summary for %s: %v", link.URL, err)
					}
				}

				c.db.UpdateLinkStatus(link.URL, "completed", nil)

				// Add delay between requests
				time.Sleep(delay)
			}(link)
		}

		// Wait for all goroutines to complete
		wg.Wait()
		close(errChan)

		// Check for any errors
		for err := range errChan {
			if err != nil {
				return fmt.Errorf("error processing links: %w", err)
			}
		}
	}

	return nil
}

// collectLinksFromURL collects links from a specific URL
func (c *Crawler) collectLinksFromURL(targetURL string, currentDepth int) error {
	collector := colly.NewCollector(
		colly.AllowedDomains(c.baseURL.Host),
	)

	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link == "" {
			return
		}

		parsedLink, err := url.Parse(link)
		if err != nil || parsedLink.Host != c.baseURL.Host {
			return
		}

		if shouldIgnoreURL(link, c.ignore) {
			return
		}

		depth := currentDepth + 1
		if depth <= c.depth {
			if err := c.db.QueueLink(link, depth); err != nil {
				debugf("Error queueing new link %s: %v", link, err)
			} else {
				debugf("Queued new link: %s (depth: %d)", link, depth)
			}
		}
	})

	return collector.Visit(targetURL)
}

// fetch retrieves content from a URL using the Reader API
func (c *Crawler) fetch(targetURL string) (string, error) {
	readerURL := fmt.Sprintf("%s/%s", strings.TrimRight(c.readerAPIURL, "/"), targetURL)
	req, err := http.NewRequest("GET", readerURL, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	debugf("Fetching content for URL: %s", targetURL)
	req.Header.Set("X-Respond-With", c.format)

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error fetching content: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}
