package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"snake-game/api/db"
)

// GetUsers to get all the user from DB
func GetUsers() http.HandlerFunc {
	data, err := db.DBConnection()
	if err != nil {
		log.Fatal("error connection with DB", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := db.GetUsersDB(data)
		if err != nil {
			log.Fatal("error getting users from the DB")
		}
		_ = json.NewEncoder(w).Encode(users)
	}
}
