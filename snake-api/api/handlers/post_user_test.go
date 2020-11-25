package handlers

import (
	"net/http"
	"snake-game/api/db"
	"snake-game/api/mock_test"
	"snake-game/api/model"
	"testing"
)

func TestPostMessage(t *testing.T) {
	usr := model.NewUser()

	headers := http.Header{}
	headers.Add("content-type", "application/json")

	w := &mock_test.ResponseWriter{}
	r := &http.Request{
		Header: headers,
	}

	r.Body = mock_test.RequestBody(map[string]string{
		"user_id":  "12345",
		"username": "Melkin",
		"score":    "500",
	})

	handler := db.CreateUser()
	handler(w, r)

	result := w.GetBodyString()

	if result != "Great job!!" {
		t.Errorf("Handler did not complete")
	}
	if len(usr.GetAll()) != 1 {
		t.Errorf("User did not add")
	}

	if usr.GetAll()[0].Username != "Melkin" {
		t.Errorf("Wrong username")
	}

	if usr.GetAll()[0].Score != "500" {
		t.Errorf("Bad score")
	}
}
