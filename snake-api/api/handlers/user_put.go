package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"snake-game/api/db"
)

// PutUser to update the score of a user by id
func PutUser() http.HandlerFunc {
	data, err := db.DBConnection()
	if err != nil {
		log.Fatal("error connecting with DB")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		_ = json.NewDecoder(r.Body).Decode(&request)

		users, err := db.GetUsersDB(data)
		if err != nil {
			log.Fatal("error getting users from DB")
		}

		for _, user := range users {
			userID := chi.URLParam(r, "id")
			if user.UserID == userID {
				user.Score = request["score"]

				db.PutUser(user, data)

				_, _ = w.Write([]byte("Great job!!"))
			}
		}
	}
}
