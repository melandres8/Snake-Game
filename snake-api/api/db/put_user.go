package db

import (
	"database/sql"
	"errors"
	"snake-game/api/model"
)

var (
	// ErrUpdatingUser to make an update of the score
	ErrUpdatingUser = errors.New("error updating user score")
)

// PutUser to update the score of the user in the DB
func PutUser(user model.User, db *sql.DB) (bool, error) {
	if _, err := db.Exec(
		"UPDATE user_tb SET score = ('" + user.Score + "') WHERE id = '" + user.UserID + "'"); err != nil {
		return false, ErrUpdatingUser
	}

	return true, nil
}
