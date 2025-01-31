package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// URLStatus represents the current state of a URL in the crawling process
type URLStatus string

const (
	StatusPending    URLStatus = "pending"
	StatusComplete   URLStatus = "complete"
	StatusFailed     URLStatus = "failed"
	StatusSkipped    URLStatus = "skipped"
	StatusProcessing URLStatus = "processing"
)

// URLRecord represents a URL and its associated metadata in the database
type URLRecord struct {
	URL          string
	Status       URLStatus
	LastAttempt  time.Time
	Attempts     int
	ErrorMessage string
	ContentPath  string
	MetadataPath string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Metadata contains additional information about the stored content
type Metadata struct {
	URL        string            `json:"url"`
	Title      string            `json:"title"`
	Timestamp  time.Time         `json:"timestamp"`
	Format     string            `json:"format"`
	Properties map[string]string `json:"properties"`
}

// Storage defines the interface for content storage operations
type Storage interface {
	// Initialize sets up the storage system
	Initialize() error

	// StoreContent saves the content and its metadata
	StoreContent(url string, content string, metadata *Metadata) error

	// GetContent retrieves stored content by URL
	GetContent(url string) (string, *Metadata, error)

	// UpdateURLStatus updates the status of a URL in the database
	UpdateURLStatus(url string, status URLStatus, errorMsg string) error

	// GetPendingURLs retrieves URLs that need to be processed
	GetPendingURLs(limit int) ([]string, error)

	// Cleanup removes all existing content and resets the database
	Cleanup() error

	// Close cleans up any resources
	Close() error
}

// SQLiteStorage implements Storage using SQLite and filesystem
type SQLiteStorage struct {
	db         *sql.DB
	baseDir    string
	contentDir string
	metaDir    string
}

// NewSQLiteStorage creates a new SQLite-based storage
func NewSQLiteStorage(baseDir string) (*SQLiteStorage, error) {
	contentDir := filepath.Join(baseDir, "content")
	metaDir := filepath.Join(baseDir, "metadata")
	dbPath := filepath.Join(baseDir, "crawler.db")

	// Create directories if they don't exist
	for _, dir := range []string{baseDir, contentDir, metaDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return &SQLiteStorage{
		db:         db,
		baseDir:    baseDir,
		contentDir: contentDir,
		metaDir:    metaDir,
	}, nil
}

// Initialize creates the necessary database tables
func (s *SQLiteStorage) Initialize() error {
	schema := `
	CREATE TABLE IF NOT EXISTS urls (
		url TEXT PRIMARY KEY,
		status TEXT NOT NULL,
		last_attempt DATETIME,
		attempts INTEGER DEFAULT 0,
		error_message TEXT,
		content_path TEXT,
		metadata_path TEXT,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_urls_status ON urls(status);
	`

	_, err := s.db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	return nil
}

// Cleanup removes all existing content and resets the database
func (s *SQLiteStorage) Cleanup() error {
	// Remove content and metadata directories
	if err := os.RemoveAll(s.contentDir); err != nil {
		return fmt.Errorf("failed to remove content directory: %w", err)
	}
	if err := os.RemoveAll(s.metaDir); err != nil {
		return fmt.Errorf("failed to remove metadata directory: %w", err)
	}

	// Recreate directories
	for _, dir := range []string{s.contentDir, s.metaDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Initialize database schema first
	if err := s.Initialize(); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	// Clear database
	_, err := s.db.Exec("DELETE FROM urls")
	if err != nil {
		return fmt.Errorf("failed to clear database: %w", err)
	}

	return nil
}

// StoreContent saves the content and metadata to files and updates the database
func (s *SQLiteStorage) StoreContent(url string, content string, metadata *Metadata) error {
	// Generate file paths
	filename := generateFilename(url, metadata.Format)
	contentPath := filepath.Join(s.contentDir, filename)
	metadataPath := filepath.Join(s.metaDir, strings.TrimSuffix(filename, filepath.Ext(filename))+".json")

	// Save content
	if err := os.WriteFile(contentPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write content: %w", err)
	}

	// Save metadata
	metadataBytes, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w", err)
	}
	if err := os.WriteFile(metadataPath, metadataBytes, 0644); err != nil {
		return fmt.Errorf("failed to write metadata: %w", err)
	}

	// Update database
	query := `
		INSERT INTO urls (url, status, content_path, metadata_path, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
		ON CONFLICT(url) DO UPDATE SET
			status = ?,
			content_path = ?,
			metadata_path = ?,
			updated_at = ?
	`

	now := time.Now()
	_, err = s.db.Exec(query,
		url, StatusComplete, contentPath, metadataPath, now, now,
		StatusComplete, contentPath, metadataPath, now)
	if err != nil {
		return fmt.Errorf("failed to update database: %w", err)
	}

	return nil
}

// GetContent retrieves stored content and metadata by URL
func (s *SQLiteStorage) GetContent(url string) (string, *Metadata, error) {
	var contentPath, metadataPath string
	err := s.db.QueryRow("SELECT content_path, metadata_path FROM urls WHERE url = ?", url).
		Scan(&contentPath, &metadataPath)
	if err != nil {
		return "", nil, fmt.Errorf("failed to query database: %w", err)
	}

	content, err := os.ReadFile(contentPath)
	if err != nil {
		return "", nil, fmt.Errorf("failed to read content: %w", err)
	}

	metadataBytes, err := os.ReadFile(metadataPath)
	if err != nil {
		return "", nil, fmt.Errorf("failed to read metadata: %w", err)
	}

	var metadata Metadata
	if err := json.Unmarshal(metadataBytes, &metadata); err != nil {
		return "", nil, fmt.Errorf("failed to unmarshal metadata: %w", err)
	}

	return string(content), &metadata, nil
}

// UpdateURLStatus updates the status of a URL in the database
func (s *SQLiteStorage) UpdateURLStatus(url string, status URLStatus, errorMsg string) error {
	query := `
		INSERT INTO urls (url, status, error_message, attempts, last_attempt, created_at, updated_at)
		VALUES (?, ?, ?, 1, ?, ?, ?)
		ON CONFLICT(url) DO UPDATE SET
			status = ?,
			error_message = ?,
			attempts = attempts + 1,
			last_attempt = ?,
			updated_at = ?
	`

	now := time.Now()
	_, err := s.db.Exec(query,
		url, status, errorMsg, now, now, now,
		status, errorMsg, now, now)
	if err != nil {
		return fmt.Errorf("failed to update URL status: %w", err)
	}

	return nil
}

// GetPendingURLs retrieves URLs that need to be processed
func (s *SQLiteStorage) GetPendingURLs(limit int) ([]string, error) {
	rows, err := s.db.Query(`
		SELECT url FROM urls 
		WHERE status IN (?, ?) 
		ORDER BY last_attempt ASC NULLS FIRST 
		LIMIT ?`,
		StatusPending, StatusFailed, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query pending URLs: %w", err)
	}
	defer rows.Close()

	var urls []string
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err != nil {
			return nil, fmt.Errorf("failed to scan URL: %w", err)
		}
		urls = append(urls, url)
	}

	return urls, nil
}

// Close closes the database connection
func (s *SQLiteStorage) Close() error {
	return s.db.Close()
}

// generateFilename creates a human-readable filename from a URL
func generateFilename(urlStr string, format string) string {
	// Parse the URL
	u, err := url.Parse(urlStr)
	if err != nil {
		// Fallback to simple hash if URL parsing fails
		hash := 0
		for i := 0; i < len(urlStr); i++ {
			hash = 31*hash + int(urlStr[i])
		}
		return fmt.Sprintf("%x", hash)
	}

	// Clean the path
	path := u.Path
	if path == "" || path == "/" {
		path = "index"
	}

	// Replace special characters
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")
	path = strings.ReplaceAll(path, "/", "-")
	path = strings.ReplaceAll(path, " ", "-")
	path = strings.ReplaceAll(path, ".", "-")

	// Add hostname to ensure uniqueness
	filename := fmt.Sprintf("%s-%s", u.Hostname(), path)

	// Ensure the filename isn't too long
	if len(filename) > 200 {
		filename = filename[:200]
	}

	// Add the appropriate extension
	switch format {
	case "text":
		filename += ".txt"
	case "html":
		filename += ".html"
	default:
		filename += ".md"
	}

	return filename
}
