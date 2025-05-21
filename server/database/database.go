package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// DB is a global variable for the SQLite database connection
var DB *sql.DB

// initDB initializes the SQLite database and creates the todos table if it doesn't exist
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./app.db") // Open a connection to the SQLite database file named app.db
	if err != nil {
		log.Fatal(err) // Log an error and stop the program if the database can't be opened
	}

	// SQL statement to create the todos table if it doesn't exist
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS servers (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		url TEXT
	);`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt) // Log an error if table creation fails
	}
}
