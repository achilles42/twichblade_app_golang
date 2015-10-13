package twichblade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	username := "praveen"
	password := "password"
	user := new(User)
	result := user.Register()
	assert.Equal(t, true, result)
	user.NewUser(username, password)
	conn, _ := user.conn.Connect()
	conn.Query("delete from users")
}

func TestUsernameDoesNotExists(t *testing.T) {
	username := "praveen"
	password := "password"
	user := new(User)
	user.NewUser(username, password)
	result := user.UsernameExists()
	assert.Equal(t, false, result)
}

func TestusernameExists(t *testing.T) {
	username := "praveen"
	password := "password"
	user := new(User)
	user.NewUser(username, password)
	conn, _ := user.conn.Connect()
	conn.Query("insert into users (username, password) values($1, $2)", username, password)
	result := user.UsernameExists()
	assert.Equal(t, true, result)
	conn.Query("delete from users")
}

func TestLoginForRegistredUser(t *testing.T) {
	username := "praveen"
	password := "password"
	user := new(User)
	user.NewUser(username, password)
	conn, _ := user.conn.Connect()
	conn.Query("insert into users(username, password) values($1, $2)", username, password)
	result := user.Login()
	assert.Equal(t, true, result)
	conn.Query("delete from users")
}

func TestLoginForNonRegistredUser(t *testing.T) {
	username := "praveen"
	password := "password"
	username1 := "praveen1"
	password1 := "password1"
	user := new(User)
	user.NewUser(username, password)
	conn, _ := user.conn.Connect()
	conn.Query("insert into users(username, password) values($1, $2)", username1, password1)
	result := user.Login()
	assert.Equal(t, false, result)
	conn.Query("delete from users")
}
