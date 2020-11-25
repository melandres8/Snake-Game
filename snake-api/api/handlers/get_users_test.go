package handlers

import (
	"net/http"
	"snake-game/api/mock_test"
	"snake-game/api/model"
	"testing"
)

func TestGetMessages(t *testing.T) {
	usr := model.NewUser()
	usr.Add(model.User{
		UserID:   "12345",
		Username: "Melkin",
		Score:    "500",
	})

	handler := GetUsers()

	w := &mock_test.ResponseWriter{}
	r := &http.Request{}

	handler(w, r)

	result := w.GetBodyJSONArray()

	if len(result) != 1 {
		t.Errorf("User was not added to the datastore")
	}

	if result[0]["user_id"] != "12345" {
		t.Errorf("User was not propely set")
	}
}
