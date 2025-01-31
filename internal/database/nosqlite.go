//go:build nosqlite
// +build nosqlite

package database

import "fmt"

// New creates a new database instance
func New(path string) (*DB, error) {
	return nil, fmt.Errorf("SQLite support is not enabled in this build")
}

// DB represents a database connection
type DB struct{}
