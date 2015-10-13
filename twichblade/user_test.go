package twichblade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistration(t *testing.T) {
	username := "praveen"
	password := "password"
	user := new(User)
	result := user.Registration(username, password)
	assert.Equal(t, true, result)
	user.NewConnection()
	conn, _ := user.conn.Connect()
	conn.Query("delete from users")
}

func TestUsernameDoesNotExists(t *testing.T) {
	username := "praveen"
	user := new(User)
	result := user.UsernameExists(username)
	assert.Equal(t, false, result)
}

func TestusernameExists(t *testing.T) {
	username := "praveen"
	password := "password"
	user := new(User)
	user.NewConnection()
	conn, _ := user.conn.Connect()
	conn.Query("insert into users (username, password) values($1, $2)", username, password)
	result := user.UsernameExists(username)
	assert.Equal(t, true, result)
	conn.Query("delete from users")

}
