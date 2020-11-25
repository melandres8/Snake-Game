package db

import (
	"database/sql"
	"fmt"
	"snake-game/api/model"
)

// PutUser to update the score of the user in the DB
func PutUser(user model.User, db *sql.DB) bool {
	if _, err := db.Exec(
		"UPDATE user_tb SET score = ('" + user.Score + "') WHERE id = '" + user.UserID + "'"); err != nil {
		fmt.Println("Error updating user score", err)
		return false
	}

	return true
}
