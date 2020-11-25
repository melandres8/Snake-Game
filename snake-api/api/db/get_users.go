package db

import (
	"database/sql"
	"fmt"
	"snake-game/api/model"
	"time"
)

// GetUsersDB to get all the user fro the DB
func GetUsersDB(db *sql.DB) []model.User {
	rows, err := db.Query("SELECT * FROM user_tb ORDER BY score DESC, created_at DESC")
	if err != nil {
		fmt.Println("Error selecting user", err)
	}

	defer rows.Close()

	var user model.User
	var id, username, score string
	var createdAt *time.Time
	users := make([]model.User, 0)

	for rows.Next() {
		if err := rows.Scan(&id, &username, &score, &createdAt); err != nil {
			fmt.Println("Error getting data", err)
		}
		user.UserID = id
		user.Username = username
		user.Score = score
		user.CreatedAt = createdAt
		users = append(users, user)
	}

	return users
}
