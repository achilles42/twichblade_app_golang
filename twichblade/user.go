package twichblade

import (
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	conn     *DbConnection
	username string
	password string
}

func (user *User) NewConnection(username, password string) {
	user.conn = new(DbConnection)
	user.username = username
	user.password = password
}

func (user *User) UsernameExists() bool {
	conn, err := user.conn.Connect()
	if err != nil {
		log.Fatal(err)
		return true
	} else {
		result, _ := conn.Query("select username from users where username = $1", user.username)
		conn.Close()
		if result.Next() {
			return true
		} else {
			return false
		}
	}
}
func (user *User) Register() bool {
	conn, err := user.conn.Connect()
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		isPresent := user.UsernameExists()
		if isPresent == false {
			_, err := conn.Query("insert into users(username, password) values($1, $2)", user.username, user.password)
			conn.Close()
			if err != nil {
				return false
			} else {
				return true
			}
		} else {
			return false
		}
	}
}

func (user *User) Login() bool {
	conn, err := user.conn.Connect()
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		result, _ := conn.Query("select (username, password) from users where username = $1 and password = $2", user.username, user.password)
		if result.Next() {
			return true
		} else {
			return false
		}
	}
}
