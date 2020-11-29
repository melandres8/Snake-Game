package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"snake-game/api/db"
	"snake-game/api/model"
)

// PostUser to post a user in a DB
func PostUser() http.HandlerFunc {
	data, err := db.DBConnection()
	if err != nil {
		log.Fatal("error connecting with DB", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		_ = json.NewDecoder(r.Body).Decode(&request)

		users, err := db.GetUsersDB(data)
		if err != nil {
			log.Fatal("error getting users from DB", err)
		}
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

		_, err = db.CreateUser(info, data)
		if err != nil {
			log.Fatal("error creating a user", err)
		}

		_, err = w.Write([]byte("Great job!!"))
		if err != nil {
			log.Fatal("error with the request", err)
		}
	}
}
