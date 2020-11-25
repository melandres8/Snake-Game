package model

import (
	"testing"
)

func TestUserGame_Add(t *testing.T) {
	usr := NewUser()

	usr.Add(User{
		"12345",
		"Melkin",
		"500",
		"12",
	})
	if len(usr.Users) != 1 {
		t.Errorf("User was not added")
	}
}

func TestUserGame_GetAll(t *testing.T) {
	usr := NewUser()
	usr.Add(User{
		"12345",
		"Melkin",
		"500",
		"12"
	})
	messages := usr.GetAll()
	if len(messages) != 1 {
		t.Errorf("There is no messages here!")
	}
}
