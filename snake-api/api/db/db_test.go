package db

import (
	"github.com/stretchr/testify/require"
	"snake-game/api/model"
	"testing"
)

func TestDBConnection(t *testing.T) {
	c := require.New(t)

	dbConnect, err := DBConnection()
	c.NoError(err)
	_ = dbConnect.Close()
}

func TestCreateUser(t *testing.T) {
	c := require.New(t)

	dbConnect, err := DBConnection()
	c.NoError(err)

	defer dbConnect.Close()

	user := model.User{
		UserID:   "123455",
		Username: "testing",
		Score:    "200",
	}

	_, err = CreateUser(user, dbConnect)
	c.NoError(err)
	_, err = dbConnect.Exec("DELETE FROM user_tb WHERE username = 'testing'")
	c.NoError(err)
}

func TestGetUsersDB(t *testing.T) {
	c := require.New(t)

	dbConnect, err := DBConnection()
	c.NoError(err)

	defer dbConnect.Close()

	_, err = dbConnect.Exec("INSERT INTO user_tb (id, username, score) VALUES ('12345', 'testing', '450')")
	c.NoError(err)
	_, err = GetUsersDB(dbConnect)
	c.NoError(err)
	rows, err := dbConnect.Query("SELECT * FROM user_tb")
	c.NoError(err)

	var users []model.User
	user := model.User{}

	for rows.Next() {
		err = rows.Scan(&user.UserID, &user.Username, &user.Score, &user.CreatedAt)
		c.NoError(err)

		users = append(users, user)
	}

	c.Equal("12345", "12345")
	_, err = dbConnect.Exec("DELETE FROM user_tb WHERE username = 'testing'")
	c.NoError(err)
}

func TestPutUser(t *testing.T) {
	c := require.New(t)

	dbConnect, err := DBConnection()
	c.NoError(err)

	defer dbConnect.Close()

	_, err = dbConnect.Exec("INSERT INTO user_tb (id, username, score) VALUES ('12345', 'testing', '450')")
	c.NoError(err)

	user := model.User{
		UserID:   "12345",
		Username: "testing",
		Score:    "450",
	}

	_, err = PutUser(user, dbConnect)
	c.NoError(err)

	var score string

	row := dbConnect.QueryRow("SELECT score FROM user_tb WHERE id = '12345'")
	err = row.Scan(&score)
	c.NoError(err)

	c.Equal(score, user.Score)
	c.Equal("12345", "12345")
	_, err = dbConnect.Exec("DELETE FROM user_tb WHERE username = 'testing'")
	c.NoError(err)
}
