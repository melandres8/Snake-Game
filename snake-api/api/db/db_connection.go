package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	// ErrDBConnection to handle the connection error from db
	ErrDBConnection = errors.New("error connection with db")
	// ErrCreateTable to handle the creation of a table in the db
	ErrCreateTable = errors.New("error creating table")
)

// DBConnection to connect with the db
func DBConnection() (*sql.DB, error) {
	// Connect to the database
	db, err := sql.Open("postgres", "postgresql://testuser@localhost:26258/users?sslmode=disable")
	if err != nil {
		return nil, ErrDBConnection
	}
	fmt.Println("Database users opened and ready.")
	fmt.Println()

	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS user_tb (id VARCHAR(100) PRIMARY KEY, username VARCHAR(50), score VARCHAR(50), created_at TIMESTAMP default now())"); err != nil {
		return nil, ErrCreateTable
	}

	return db, nil
}
