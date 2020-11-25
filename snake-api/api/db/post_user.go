package db

import (
	"database/sql"
	"fmt"
	"snake-game/api/model"
)

// CreateUser to post a new user
func CreateUser(user model.User, db *sql.DB) bool {
	// Insert a new user in user_tb
	if user.UserID != "" {
		_, err := db.Exec(
			"INSERT INTO user_tb (id, username, score) VALUES ('" + user.UserID + "', '" + user.Username + "', '" + user.Score + "')",
		)
		if err != nil {
			fmt.Println("Error inserting data", err)
			return false
		}
	}

	return true
}
