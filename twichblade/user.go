package twichblade

import (
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	conn *DbConnection
}

func (user *User) NewConnection() {
	user.conn = new(DbConnection)
}

func (user *User) UsernameExists(username string) bool {
	conn, err := user.conn.Connect()
	if err != nil {
		log.Fatal(err)

		return true
	} else {
		result, _ := conn.Query("select username from users where username = $1", username)
		conn.Close()
		if result.Next() {
			return true
		} else {
			return false
		}
	}
}

func (user *User) Registration(username, password string) bool {
	conn, err := user.conn.Connect()
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		isPresent := user.UsernameExists(username)
		if isPresent == false {
			_, err := conn.Query("insert into users(username, password) values($1, $2)", username, password)
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
