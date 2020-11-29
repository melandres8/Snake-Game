package db

import (
	"database/sql"
	"errors"
	"snake-game/api/model"
	"time"
)

var (
	// ErrSelectionUsers to get all user from the db
	ErrSelectionUsers = errors.New("error selecting user")
	// ErrGettingData to get all data from the db
	ErrGettingData = errors.New("error getting data")
)

// GetUsersDB to get all the user fro the DB
func GetUsersDB(db *sql.DB) ([]model.User, error) {
	rows, err := db.Query("SELECT * FROM user_tb ORDER BY score DESC, created_at DESC")
	if err != nil {
		return nil, ErrSelectionUsers
	}

	defer rows.Close()

	var user model.User
	var id, username, score string
	var createdAt time.Time
	users := make([]model.User, 0)

	for rows.Next() {
		if err := rows.Scan(&id, &username, &score, &createdAt); err != nil {
			return nil, ErrGettingData
		}
		user.UserID = id
		user.Username = username
		user.Score = score
		user.CreatedAt = createdAt
		users = append(users, user)
	}

	return users, nil
}
