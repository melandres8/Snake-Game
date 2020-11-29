package model

import (
	"github.com/google/uuid"
	"time"
)

// User model
type User struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Score     string    `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}

// GenerateID randomly for any user
func GenerateID() string {
	return uuid.New().String()
}
