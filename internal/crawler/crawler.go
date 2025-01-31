package crawler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"
	"time"

	"stripper/internal/database"
	"stripper/internal/storage"
	"stripper/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gocolly/colly/v2"
)

// Crawler handles the web crawling functionality
type Crawler struct {
	client    *http.Client
	baseURL   *url.URL
	depth     int
	format    string
	force     bool
	ignore    []string
	outputDir string
	storage   storage.Storage
	db        *database.DB
	ui        *tea.Program
}

// Options configures the crawler behavior
type Options struct {
	URL       string
	Depth     int
	Format    string
	Force     bool
	Ignore    []string
	OutputDir string
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

	// Create crawler instance
	c := &Crawler{
		client:    &http.Client{},
		baseURL:   baseURL,
		depth:     opts.Depth,
		format:    opts.Format,
		force:     opts.Force,
		ignore:    opts.Ignore,
		outputDir: opts.OutputDir,
		storage:   store,
		db:        db,
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
	collector := colly.NewCollector(
		colly.MaxDepth(c.depth),
		colly.AllowedDomains(c.baseURL.Host),
		colly.Async(true),
	)

	// Add rate limiting
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       1 * time.Second,
	})

	debugf("Starting link collection for %s with depth %d", c.baseURL.String(), c.depth)

	// Queue the initial URL
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

		depth := e.Request.Depth + 1
		if err := c.db.QueueLink(link, depth); err != nil {
			debugf("Error queueing link %s: %v", link, err)
		} else {
			debugf("Queued link: %s (depth: %d)", link, depth)
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
		batchSize = 10
		delay     = 500 * time.Millisecond
	)

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

		// Process each link
		for _, link := range links {
			debugf("Processing link: %s (depth: %d)", link.URL, link.Depth)

			// Check if we should recrawl
			shouldCrawl, err := c.db.ShouldRecrawl(link.URL, c.force, 24*time.Hour)
			if err != nil {
				c.db.UpdateLinkStatus(link.URL, "failed", err)
				continue
			}

			if !shouldCrawl {
				debugf("Skipping recent URL: %s (last crawled: %s)", link.URL, link.LastCrawled)
				continue
			}

			debugf("Fetching content for URL: %s", link.URL)

			// Fetch content using Reader API
			content, err := c.fetch(link.URL)
			if err != nil {
				c.db.UpdateLinkStatus(link.URL, "failed", err)
				continue
			}

			// Store content
			if err := c.storage.Save(link.URL, content, c.format); err != nil {
				c.db.UpdateLinkStatus(link.URL, "failed", err)
				continue
			}

			c.db.UpdateLinkStatus(link.URL, "completed", nil)

			// Add delay between requests
			time.Sleep(delay)
		}
	}

	return nil
}

// fetch retrieves content from a URL using the Reader API
func (c *Crawler) fetch(targetURL string) (string, error) {
	readerURL := fmt.Sprintf("https://read.tabnot.space/%s", targetURL)
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

// shouldIgnoreURL checks if a URL should be ignored based on its extension
func shouldIgnoreURL(urlStr string, ignoreExts []string) bool {
	lower := strings.ToLower(urlStr)
	for _, ext := range ignoreExts {
		if strings.HasSuffix(lower, "."+strings.TrimPrefix(ext, ".")) {
			return true
		}
	}
	return false
}
