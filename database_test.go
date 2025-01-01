package main

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestInitDatabase(t *testing.T) {
	// Call InitDatabase to initialize the database connection and create the table
	InitDatabase()
	defer DB.Close() // Ensure the database connection is closed after the test

	// Query to check if the "users" table exists
	query := `SELECT name FROM sqlite_master WHERE type='table' AND name='users';`
	var tableName string
	err := DB.QueryRow(query).Scan(&tableName)
	if err != nil {
		// Fail the test if the table does not exist
		t.Fatalf("Error checking table existence: %v", err)
	}

	// Verify the table name matches "users"
	if tableName != "users" {
		t.Fatalf("Expected table 'users' but got '%s'", tableName)
	}

	// Log success
	t.Log("InitDatabase successfully created the 'users' table.")
}
