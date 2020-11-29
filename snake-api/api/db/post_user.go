package db

import (
	"database/sql"
	"errors"
	"snake-game/api/model"
)

var (
	// ErrInsertingData to post a user in a DB
	ErrInsertingData = errors.New("error inserting data")
)

// CreateUser to post a new user
func CreateUser(user model.User, db *sql.DB) (bool, error) {
	// Insert a new user in user_tb
	if user.UserID != "" {
		_, err := db.Exec(
			"INSERT INTO user_tb (id, username, score) VALUES ('" + user.UserID + "', '" + user.Username + "', '" + user.Score + "')",
		)
		if err != nil {
			return false, ErrInsertingData
		}
	}

	return true, nil
}
