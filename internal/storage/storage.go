package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Storage defines the interface for storing crawled content
type Storage interface {
	Save(url string, content string, format string) error
	HasContent(url string) bool
	GetLastCrawled(url string) (time.Time, error)
}

// FileStorage implements Storage using the local filesystem
type FileStorage struct {
	baseDir string
}

// NewFileStorage creates a new FileStorage instance
func NewFileStorage(baseDir string) (*FileStorage, error) {
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}
	return &FileStorage{baseDir: baseDir}, nil
}

// Save stores the content to a file
func (fs *FileStorage) Save(url string, content string, format string) error {
	filename := fs.urlToFilename(url, format)
	fullPath := filepath.Join(fs.baseDir, filename)

	// Create metadata
	metadata := fmt.Sprintf("URL: %s\nDate: %s\n\n", url, time.Now().Format(time.RFC3339))
	content = metadata + content

	// Write content to file
	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// HasContent checks if content exists for the given URL
func (fs *FileStorage) HasContent(url string) bool {
	formats := []string{"markdown", "text", "html"}
	for _, format := range formats {
		filename := fs.urlToFilename(url, format)
		if _, err := os.Stat(filepath.Join(fs.baseDir, filename)); err == nil {
			return true
		}
	}
	return false
}

// GetLastCrawled returns the last crawl time for the given URL
func (fs *FileStorage) GetLastCrawled(url string) (time.Time, error) {
	formats := []string{"markdown", "text", "html"}
	var lastMod time.Time

	for _, format := range formats {
		filename := fs.urlToFilename(url, format)
		info, err := os.Stat(filepath.Join(fs.baseDir, filename))
		if err == nil && (lastMod.IsZero() || info.ModTime().After(lastMod)) {
			lastMod = info.ModTime()
		}
	}

	if lastMod.IsZero() {
		return lastMod, fmt.Errorf("no content found for URL: %s", url)
	}

	return lastMod, nil
}

// urlToFilename converts a URL to a safe filename
func (fs *FileStorage) urlToFilename(url string, format string) string {
	// Remove scheme and query parameters
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")
	if idx := strings.Index(url, "?"); idx != -1 {
		url = url[:idx]
	}

	// Replace special characters
	url = strings.ReplaceAll(url, "/", "_")
	url = strings.ReplaceAll(url, ":", "_")

	// Add extension based on format
	ext := ".md"
	switch format {
	case "text":
		ext = ".txt"
	case "html":
		ext = ".html"
	}

	return url + ext
}
