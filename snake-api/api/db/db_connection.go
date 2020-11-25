package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DBConnection to connect with the db
func DBConnection() *sql.DB {
	// Connect to the database
	db, err := sql.Open("postgres", "postgresql://testuser@localhost:26258/users?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database ", err)
	}
	fmt.Println("Database users opened and ready.")
	fmt.Println()

	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS user_tb (id VARCHAR(100) PRIMARY KEY, username VARCHAR(50), score VARCHAR(50), created_at TIMESTAMP default now())"); err != nil {
		fmt.Println("Error creating table: ", err)
	}

	return db
}
