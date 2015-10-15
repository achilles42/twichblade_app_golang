package twichblade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	username := "foo"
	password := "bar"
	user := new(User)
	user.NewUser(username, password)
	result, err := user.Register()
	assert.Equal(t, true, result)
	assert.Equal(t, nil, err)
	conn, _ := user.conn.Connect()
	conn.Query("delete from users")
}

func TestUsernameDoesNotExists(t *testing.T) {
	username := "foo"
	password := "bar"
	user := new(User)
	user.NewUser(username, password)
	result, err := user.UsernameExists()
	assert.Equal(t, false, result)
	assert.Equal(t, nil, err)
}

func TestUsernameExists(t *testing.T) {
	username := "foo4"
	password := "bar4"
	user := new(User)
	user.NewUser(username, password)
	conn, _ := user.conn.Connect()
	conn.QueryRow("insert into users (username, password) values($1, $2)", username, password)
	result, err := user.UsernameExists()
	assert.Equal(t, true, result)
	assert.Equal(t, nil, err)
	conn.Query("delete from users")
}

func TestLoginForRegistredUser(t *testing.T) {
	username := "foo1"
	password := "bar1"
	user := new(User)
	user.NewUser(username, password)
	conn, _ := user.conn.Connect()
	conn.QueryRow("insert into users(username, password) values($1, $2)", username, password)
	result, err := user.Login()
	assert.Equal(t, true, result)
	assert.Equal(t, nil, err)
	conn.Query("delete from users")
}

func TestLoginForNonRegistredUser(t *testing.T) {
	username := "foo"
	password := "bar"
	username1 := "foo1"
	password1 := "bar1"
	user := new(User)
	user.NewUser(username, password)
	conn, _ := user.conn.Connect()
	conn.Query("insert into users(username, password) values($1, $2)", username1, password1)
	result, err := user.Login()
	assert.Equal(t, false, result)
	assert.Equal(t, nil, err)
	conn.Query("delete from users")
}
