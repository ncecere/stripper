package crawler

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/ncecere/stripper/internal/storage"
	"github.com/ncecere/stripper/pkg/reader"
)

var debugLog = log.New(log.Writer(), "[DEBUG] ", 0)

func init() {
	// Disable debug logging by default
	debugLog.SetOutput(io.Discard)
}

// crawlerImpl implements the Crawler interface
type crawlerImpl struct {
	config     Config
	reader     *reader.Client
	storage    storage.Storage
	stats      Stats
	statsMu    sync.RWMutex
	queue      chan urlDepth
	done       chan struct{}
	callback   func(Stats)
	wg         sync.WaitGroup
	urlsSeen   map[string]struct{}
	urlSeenMu  sync.RWMutex
	processing sync.WaitGroup
}

type urlDepth struct {
	url   string
	depth int
}

// NewCrawler creates a new crawler instance
func NewCrawler(readerClient *reader.Client, storage storage.Storage) Crawler {
	return &crawlerImpl{
		reader:  readerClient,
		storage: storage,
		stats: Stats{
			ErrorsByType: make(map[string]int),
			StartTime:    time.Now(),
		},
		queue:    make(chan urlDepth, 1000),
		urlsSeen: make(map[string]struct{}),
		done:     make(chan struct{}),
	}
}

// Start begins the crawling process
func (c *crawlerImpl) Start(ctx context.Context, config Config) error {
	c.config = config

	debugLog.Printf("Starting crawler with config:")
	debugLog.Printf("- Base URL: %s", config.BaseURL)
	debugLog.Printf("- Max Depth: %d", config.MaxDepth)
	debugLog.Printf("- Stay in Domain: %v", config.StayInDomain)
	debugLog.Printf("- Format: %s", config.Format)

	// Clean up and initialize storage
	if err := c.storage.Cleanup(); err != nil {
		return fmt.Errorf("failed to cleanup storage: %w", err)
	}
	if err := c.storage.Initialize(); err != nil {
		return fmt.Errorf("failed to initialize storage: %w", err)
	}

	debugLog.Printf("Starting link collection phase...")
	// First phase: collect all links
	if err := c.collectLinks(ctx, c.config.BaseURL, 0); err != nil {
		return fmt.Errorf("failed to collect links: %w", err)
	}

	// Store total URLs to process for progress calculation
	c.statsMu.Lock()
	c.stats.TotalURLs = len(c.urlsSeen)
	c.statsMu.Unlock()

	debugLog.Printf("Found %d unique URLs to process", c.stats.TotalURLs)

	// Start worker pool
	debugLog.Printf("Starting %d workers for content processing...", c.config.BatchSize)
	for i := 0; i < c.config.BatchSize; i++ {
		c.wg.Add(1)
		go c.worker(ctx)
	}

	// Queue all collected URLs for processing
	debugLog.Printf("Queueing URLs for processing...")
	for url := range c.urlsSeen {
		select {
		case c.queue <- urlDepth{url: url, depth: 0}:
			c.processing.Add(1)
		default:
			debugLog.Printf("Warning: Queue full, skipping URL: %s", url)
		}
	}

	debugLog.Printf("Waiting for processing to complete...")
	// Wait for all processing to complete
	c.processing.Wait()
	debugLog.Printf("Processing complete!")

	return nil
}

// Stop gracefully stops the crawling process
func (c *crawlerImpl) Stop() error {
	debugLog.Printf("Stopping crawler...")
	close(c.done)
	c.wg.Wait()
	return nil
}

// GetStats returns the current crawling statistics
func (c *crawlerImpl) GetStats() Stats {
	c.statsMu.RLock()
	defer c.statsMu.RUnlock()
	return c.stats
}

// AddURLs adds new URLs to the crawling queue at depth 0
func (c *crawlerImpl) AddURLs(urls []string) error {
	for _, url := range urls {
		if err := c.addURL(url, 0); err != nil {
			return err
		}
	}
	return nil
}

// addURL adds a single URL to the queue if it hasn't been seen before
func (c *crawlerImpl) addURL(url string, depth int) error {
	// Check if we should stay in domain
	if c.config.StayInDomain && !isSameDomain(c.config.BaseURL, url) {
		debugLog.Printf("Skipping out-of-domain URL: %s", url)
		return nil
	}

	// Check depth limit
	if c.config.MaxDepth > 0 && depth > c.config.MaxDepth {
		debugLog.Printf("Skipping URL due to depth limit: %s (depth: %d)", url, depth)
		return nil
	}

	// Check if URL has been seen
	c.urlSeenMu.Lock()
	if _, seen := c.urlsSeen[url]; seen {
		c.urlSeenMu.Unlock()
		debugLog.Printf("Skipping already seen URL: %s", url)
		return nil
	}
	c.urlsSeen[url] = struct{}{}
	debugLog.Printf("Found new URL: %s (depth: %d)", url, depth)
	c.urlSeenMu.Unlock()

	// Update URL status in storage
	if err := c.storage.UpdateURLStatus(url, storage.StatusPending, ""); err != nil {
		return fmt.Errorf("failed to update URL status: %w", err)
	}

	return nil
}

// SetProgressCallback sets the progress callback function
func (c *crawlerImpl) SetProgressCallback(callback func(Stats)) {
	c.callback = callback
}

// worker processes URLs from the queue
func (c *crawlerImpl) worker(ctx context.Context) {
	defer c.wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case <-c.done:
			return
		case ud := <-c.queue:
			debugLog.Printf("Worker processing URL: %s", ud.url)
			c.processURL(ctx, ud.url, ud.depth)
			// Respect rate limiting
			time.Sleep(c.config.RequestDelay)
		}
	}
}

// collectLinks recursively collects all links starting from the given URL
func (c *crawlerImpl) collectLinks(ctx context.Context, url string, depth int) error {
	debugLog.Printf("Collecting links from: %s (depth: %d)", url, depth)

	// Check depth limit
	if c.config.MaxDepth > 0 && depth > c.config.MaxDepth {
		debugLog.Printf("Skipping %s: exceeded max depth %d", url, c.config.MaxDepth)
		return nil
	}

	// Check if URL has been seen
	c.urlSeenMu.Lock()
	if _, seen := c.urlsSeen[url]; seen {
		c.urlSeenMu.Unlock()
		debugLog.Printf("Skipping %s: already seen", url)
		return nil
	}
	c.urlsSeen[url] = struct{}{}
	debugLog.Printf("Found new URL: %s (depth: %d)", url, depth)
	c.urlSeenMu.Unlock()

	// Add URL to database
	if err := c.storage.UpdateURLStatus(url, storage.StatusPending, ""); err != nil {
		return fmt.Errorf("failed to update URL status: %w", err)
	}

	// Fetch content for link extraction
	debugLog.Printf("Fetching content from: %s", url)
	content, err := c.reader.FetchContent(url, reader.FormatHTML)
	if err != nil {
		debugLog.Printf("Error fetching content from %s: %v", url, err)
		return fmt.Errorf("failed to fetch content: %w", err)
	}

	// Extract and process links
	links, err := extractLinks(content, url)
	if err != nil {
		debugLog.Printf("Error extracting links from %s: %v", url, err)
		return fmt.Errorf("failed to extract links: %w", err)
	}
	debugLog.Printf("Found %d links in %s", len(links), url)

	// Process each link
	for _, link := range links {
		// Skip if not in same domain when required
		if c.config.StayInDomain && !isSameDomain(c.config.BaseURL, link) {
			debugLog.Printf("Skipping out-of-domain link: %s", link)
			continue
		}

		// Recursively collect links
		if err := c.collectLinks(ctx, link, depth+1); err != nil {
			debugLog.Printf("Warning: failed to collect links from %s: %v", link, err)
		}
	}

	return nil
}

// processURL handles the crawling of a single URL
func (c *crawlerImpl) processURL(ctx context.Context, url string, depth int) {
	defer c.processing.Done()
	start := time.Now()

	debugLog.Printf("Processing URL: %s", url)

	// Update status to processing
	if err := c.storage.UpdateURLStatus(url, storage.StatusProcessing, ""); err != nil {
		debugLog.Printf("Error updating status for %s: %v", url, err)
		c.updateStats(url, err, 0)
		return
	}

	// Fetch content in desired format
	debugLog.Printf("Fetching content from %s in %s format", url, c.config.Format)
	content, err := c.reader.FetchContent(url, c.config.Format)
	if err != nil {
		debugLog.Printf("Error fetching content from %s: %v", url, err)
		c.updateStats(url, err, time.Since(start))
		return
	}

	// Extract title from content
	title := extractTitle(content)
	debugLog.Printf("Title extracted from %s: %s", url, title)

	// Create metadata
	metadata := &storage.Metadata{
		URL:       url,
		Title:     title,
		Timestamp: time.Now(),
		Format:    string(c.config.Format),
		Properties: map[string]string{
			"latency": time.Since(start).String(),
			"depth":   fmt.Sprintf("%d", depth),
		},
	}

	// Store content and metadata
	debugLog.Printf("Storing content and metadata for %s", url)
	if err := c.storage.StoreContent(url, content, metadata); err != nil {
		debugLog.Printf("Error storing content for %s: %v", url, err)
		c.updateStats(url, err, time.Since(start))
		return
	}

	// Update statistics
	c.updateStats(url, nil, time.Since(start))
	debugLog.Printf("Successfully processed %s", url)
}

// updateStats updates the crawler statistics
func (c *crawlerImpl) updateStats(url string, err error, latency time.Duration) {
	c.statsMu.Lock()
	defer c.statsMu.Unlock()

	c.stats.URLsProcessed++
	c.stats.LastUpdateTime = time.Now()

	// Update latency statistics
	if c.stats.AverageLatency == 0 {
		c.stats.AverageLatency = latency
	} else {
		c.stats.AverageLatency = (c.stats.AverageLatency + latency) / 2
	}

	// Update progress percentage
	if c.stats.TotalURLs > 0 {
		c.stats.Progress = float64(c.stats.URLsProcessed) / float64(c.stats.TotalURLs)
		debugLog.Printf("Progress: %.1f%% (%d/%d URLs)",
			c.stats.Progress*100,
			c.stats.URLsProcessed,
			c.stats.TotalURLs)
	}

	if err != nil {
		c.stats.URLsFailed++
		// Update error statistics
		errType := "unknown"
		if crawlerErr, ok := err.(*CrawlerError); ok {
			errType = crawlerErr.Type.String()
		}
		c.stats.ErrorsByType[errType]++

		// Update storage with error
		if err := c.storage.UpdateURLStatus(url, storage.StatusFailed, err.Error()); err != nil {
			debugLog.Printf("Failed to update URL status: %v", err)
		}
	} else {
		c.stats.URLsSuccessful++
	}

	// Notify progress callback if set
	if c.callback != nil {
		c.callback(c.stats)
	}
}

// String returns a string representation of ErrorType
func (e ErrorType) String() string {
	switch e {
	case ErrorTypeNetwork:
		return "network"
	case ErrorTypeHTTP:
		return "http"
	case ErrorTypeParser:
		return "parser"
	case ErrorTypeStorage:
		return "storage"
	default:
		return "unknown"
	}
}
