package handlers

import (
	"encoding/json"
	"net/http"
	"snake-game/api/db"
	"snake-game/api/model"
)

// PostUser to post a user in a DB
func PostUser() http.HandlerFunc {
	data := db.DBConnection()
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		_ = json.NewDecoder(r.Body).Decode(&request)

		users := db.GetUsersDB(data)
		var info model.User

		info.Username = request["username"]

		if info.Username != "" {
			info.Score = request["score"]
			info.UserID = model.GenerateID()
		}

		for _, user := range users {
			if user.Username == info.Username {
				info.Username = user.Username
				info.Score = user.Score
				info.UserID = user.UserID
				info.CreatedAt = user.CreatedAt
			}
		}

		db.CreateUser(info, data)

		_, _ = w.Write([]byte("Great job!!"))
	}
}
