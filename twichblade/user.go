package twichblade

import _ "github.com/lib/pq"

type User struct {
	conn     *DbConnection
	username string
	password string
}

func (user *User) NewUser(username, password string) {
	user.conn = new(DbConnection)
	user.username = username
	user.password = password
}

func (user *User) UsernameExists() (bool, error) {
	conn, err := user.conn.Connect()
	var username string
	if err != nil {
		return false, err
	} else {
		conn.QueryRow("select username from users where username = $1", user.username).Scan(&username)
		defer conn.Close()
		if username == user.username {
			return true, nil
		} else {
			return false, nil
		}
	}
}

func (user *User) Register() (bool, error) {
	conn, err := user.conn.Connect()
	if err != nil {
		return false, err
	} else {
		isPresent, err := user.UsernameExists()
		if err != nil {
			return false, err
		} else if isPresent == false {
			_, err := conn.Query("insert into users(username, password) values($1, $2)", user.username, user.password)
			if err != nil {
				return false, nil
			} else {
				return true, nil
			}
		} else {
			return false, nil
		}
	}
}

func (user *User) Login() (bool, error) {
	conn, err := user.conn.Connect()
	var username string
	var password string
	if err != nil {
		return false, err
	} else {
		conn.QueryRow("select username, password from users where username = $1 and password = $2", user.username, user.password).Scan(&username, &password)
		if username == user.username && password == user.password {
			return true, nil
		} else {
			return false, nil
		}
	}
}
