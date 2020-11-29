package model

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserModel(t *testing.T) {
	c := require.New(t)

	user := User{
		UserID:    GenerateID(),
		Username:  "testing",
		Score:     "500",
		CreatedAt: time.Time{},
	}

	var newUser User

	response, err := json.Marshal(user)
	c.NoError(err)

	err = json.Unmarshal(response, &newUser)
	c.NoError(err)
	c.Equal(user, newUser)
}
