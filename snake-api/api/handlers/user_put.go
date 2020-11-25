package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"snake-game/api/db"
)

// PutUser to update the score of a user by id
func PutUser() http.HandlerFunc {
	data := db.DBConnection()
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		_ = json.NewDecoder(r.Body).Decode(&request)

		users := db.GetUsersDB(data)

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
