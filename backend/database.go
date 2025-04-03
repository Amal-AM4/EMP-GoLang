package backend

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the database and returns the connection
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./employees.db")
	if err != nil {
		log.Fatal("Database connection failed:", err)
		return nil, err
	}

	// Create employees table if it doesn't exist
	query := `CREATE TABLE IF NOT EXISTS employees (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER,
		position TEXT,
		salary INTEGER
	)`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Table creation failed:", err)
		return nil, err
	}

	fmt.Println("Database initialized successfully")
	return db, nil
}
