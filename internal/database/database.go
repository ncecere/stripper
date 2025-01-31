package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// DB handles database operations
type DB struct {
	db *sql.DB
}

// Link represents a URL to be crawled
type Link struct {
	URL         string
	LastCrawled time.Time
	Depth       int
	Status      string // "pending", "completed", "failed"
	Error       string // empty string for no error
}

// New creates a new database connection and initializes tables
func New(dbPath string) (*DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := initTables(db); err != nil {
		db.Close()
		return nil, err
	}

	return &DB{db: db}, nil
}

// Close closes the database connection
func (d *DB) Close() error {
	return d.db.Close()
}

// initTables creates the necessary database tables
func initTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS links (
			url TEXT PRIMARY KEY,
			last_crawled DATETIME,
			depth INTEGER,
			status TEXT,
			error TEXT
		);
		CREATE INDEX IF NOT EXISTS idx_status ON links(status);
		CREATE INDEX IF NOT EXISTS idx_last_crawled ON links(last_crawled);
	`)
	return err
}

// QueueLink adds a link to the database
func (d *DB) QueueLink(url string, depth int) error {
	_, err := d.db.Exec(`
		INSERT OR IGNORE INTO links (url, depth, status)
		VALUES (?, ?, 'pending')
	`, url, depth)
	return err
}

// GetNextBatch returns a batch of pending links
func (d *DB) GetNextBatch(batchSize int) ([]Link, error) {
	rows, err := d.db.Query(`
		SELECT 
			url,
			last_crawled,
			depth,
			status,
			COALESCE(error, '') as error
		FROM links
		WHERE status = 'pending'
		LIMIT ?
	`, batchSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []Link
	for rows.Next() {
		var link Link
		var lastCrawled sql.NullTime
		err := rows.Scan(&link.URL, &lastCrawled, &link.Depth, &link.Status, &link.Error)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		if lastCrawled.Valid {
			link.LastCrawled = lastCrawled.Time
		}
		links = append(links, link)
	}
	return links, nil
}

// UpdateLinkStatus updates the status of a link
func (d *DB) UpdateLinkStatus(url string, status string, err error) error {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	_, dbErr := d.db.Exec(`
		UPDATE links
		SET status = ?, error = ?, last_crawled = CURRENT_TIMESTAMP
		WHERE url = ?
	`, status, errMsg, url)
	return dbErr
}

// ShouldRecrawl checks if a URL should be recrawled based on last crawl time
func (d *DB) ShouldRecrawl(url string, force bool, minAge time.Duration) (bool, error) {
	if force {
		return true, nil
	}

	var lastCrawled sql.NullTime
	err := d.db.QueryRow(`
		SELECT last_crawled
		FROM links
		WHERE url = ?
	`, url).Scan(&lastCrawled)

	if err == sql.ErrNoRows {
		return true, nil
	}
	if err != nil {
		return false, err
	}

	if !lastCrawled.Valid {
		return true, nil
	}

	return time.Since(lastCrawled.Time) > minAge, nil
}

// GetStats returns crawling statistics
func (d *DB) GetStats() (total, pending, completed, failed int, err error) {
	err = d.db.QueryRow(`
		SELECT
			COUNT(*) as total,
			SUM(CASE WHEN status = 'pending' THEN 1 ELSE 0 END) as pending,
			SUM(CASE WHEN status = 'completed' THEN 1 ELSE 0 END) as completed,
			SUM(CASE WHEN status = 'failed' THEN 1 ELSE 0 END) as failed
		FROM links
	`).Scan(&total, &pending, &completed, &failed)
	return
}
