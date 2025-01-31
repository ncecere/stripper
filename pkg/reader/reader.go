package reader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Format represents the response format from the Reader API
type Format string

const (
	FormatText     Format = "text"
	FormatMarkdown Format = "markdown"
	FormatHTML     Format = "html"
)

// Config holds the configuration for the Reader client
type Config struct {
	// BaseURL is the Reader API endpoint
	BaseURL string
	// Timeout for requests
	Timeout time.Duration
	// DefaultFormat is the preferred response format
	DefaultFormat Format
	// UserAgent identifies our client
	UserAgent string
}

// Client handles interactions with the Reader API
type Client struct {
	config     Config
	httpClient *http.Client
}

// NewClient creates a new Reader API client
func NewClient(config Config) *Client {
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}
	if config.DefaultFormat == "" {
		config.DefaultFormat = FormatMarkdown
	}
	if config.BaseURL == "" {
		config.BaseURL = "https://read.tabnot.space"
	}
	if config.UserAgent == "" {
		config.UserAgent = "Stripper/1.0"
	}

	return &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

// FetchContent retrieves content from a URL using the Reader API
func (c *Client) FetchContent(targetURL string, format Format) (string, error) {
	if format == "" {
		format = c.config.DefaultFormat
	}

	// The Reader API expects the target URL as part of the path
	requestURL := fmt.Sprintf("%s/%s", c.config.BaseURL, targetURL)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", c.config.UserAgent)
	req.Header.Set("X-Respond-With", string(format))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(content), nil
}

// ValidateURL checks if a URL is valid and can be processed by the Reader API
func (c *Client) ValidateURL(targetURL string) error {
	_, err := url.Parse(targetURL)
	if err != nil {
		return fmt.Errorf("invalid URL format: %w", err)
	}
	return nil
}
