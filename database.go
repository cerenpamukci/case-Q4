package main

import (
	"database/sql" // Package for database operations
	"log"          // Package for logging errors and messages

	_ "github.com/mattn/go-sqlite3" // SQLite3 driver, imported anonymously for side effects
)

// Global variable to hold the database connection
var DB *sql.DB

// InitDatabase initializes the SQLite database and creates the necessary table
func InitDatabase() {
	var err error

	// Open a connection to the SQLite database
	// "./users.db" specifies the file path of the database; it will be created if it doesn't exist
	DB, err = sql.Open("sqlite3", "./users.db")
	if err != nil { // Check for errors while opening the database
		log.Fatal(err) // Log the error and stop execution if the connection fails
	}

	// SQL query to create the "users" table if it doesn't already exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT, -- Auto-incrementing primary key for unique IDs
		name TEXT NOT NULL,                  -- Name column, cannot be null
		email TEXT NOT NULL,                 -- Email column, cannot be null
		phone TEXT NOT NULL                  -- Phone column, cannot be null
	);`

	// Execute the SQL query to create the table
	_, err = DB.Exec(createTableQuery)
	if err != nil { // Check for errors while executing the query
		log.Fatal(err) // Log the error and stop execution if table creation fails
	}
}
