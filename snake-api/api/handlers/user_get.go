package handlers

import (
	"encoding/json"
	"net/http"
	"snake-game/api/db"
)

// GetUsers to get all the user from DB
func GetUsers() http.HandlerFunc {
	data := db.DBConnection()
	return func(w http.ResponseWriter, r *http.Request) {
		users := db.GetUsersDB(data)
		_ = json.NewEncoder(w).Encode(users)
	}
}
