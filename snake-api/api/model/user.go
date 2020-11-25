package model

import (
	"github.com/google/uuid"
	"time"
)

// User model
type User struct {
	UserID		string		`json:"user_id"`
	Username	string		`json:"username"`
	Score		string		`json:"score"`
	CreatedAt	*time.Time	`json:"created_at"`
}

// UserGame is an array of users
type UserGame struct {
	Users []User
}

// Getter interface to get all user
type Getter interface {
	GetAll() []User
}

// Setter interface to set a user
type Setter interface {
	Add(user User)
}

// NewUser create a new user
func NewUser() *UserGame {
	return &UserGame{
		Users: []User{},
	}
}

// Add a user
func (u *UserGame) Add(user User) {
	u.Users = append(u.Users, user)
}

// GetAll users
func (u *UserGame) GetAll() []User {
	return u.Users
}

// GenerateID randomly for any user
func GenerateID() string {
	return uuid.New().String()
}
