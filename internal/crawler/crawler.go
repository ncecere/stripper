package crawler

import (
	"context"
	"time"

	"github.com/ncecere/stripper/pkg/reader"
)

// Config holds the configuration for the crawler
type Config struct {
	// BatchSize is the number of URLs to process in parallel
	BatchSize int
	// Delay between requests to respect server load
	RequestDelay time.Duration
	// MaxRetries for failed requests
	MaxRetries int
	// BaseURL is the starting point for crawling
	BaseURL string
	// OutputDir is where the markdown files will be stored
	OutputDir string
	// MaxDepth is the maximum depth to crawl (0 means no limit)
	MaxDepth int
	// StayInDomain if true, only crawl URLs in the same domain
	StayInDomain bool
	// Format specifies the desired output format
	Format reader.Format
}

// Stats tracks the crawler's progress
type Stats struct {
	URLsProcessed  int
	URLsSuccessful int
	URLsFailed     int
	TotalURLs      int // Total number of URLs to process
	StartTime      time.Time
	LastUpdateTime time.Time
	ErrorsByType   map[string]int
	AverageLatency time.Duration
	Progress       float64 // Progress percentage (0.0 to 1.0)
}

// Result represents the outcome of processing a single URL
type Result struct {
	URL       string
	Content   string
	Metadata  map[string]string
	Error     error
	Timestamp time.Time
	Latency   time.Duration
}

// Crawler defines the interface for web content crawling
type Crawler interface {
	// Start begins the crawling process with the given configuration
	Start(ctx context.Context, config Config) error

	// Stop gracefully stops the crawling process
	Stop() error

	// GetStats returns the current crawling statistics
	GetStats() Stats

	// AddURLs adds new URLs to the crawling queue
	AddURLs(urls []string) error

	// SetProgressCallback sets a function to be called when progress updates occur
	SetProgressCallback(func(Stats))
}

// ErrorType represents different categories of errors that can occur during crawling
type ErrorType int

const (
	ErrorTypeNetwork ErrorType = iota
	ErrorTypeHTTP
	ErrorTypeParser
	ErrorTypeStorage
	ErrorTypeUnknown
)

// CrawlerError represents an error that occurred during crawling
type CrawlerError struct {
	Type    ErrorType
	Message string
	URL     string
	Cause   error
}

func (e *CrawlerError) Error() string {
	return e.Message
}
