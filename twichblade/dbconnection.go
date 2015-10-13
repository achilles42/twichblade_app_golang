package twichblade

import "database/sql"

type DbConnection struct {
}

func (t *DbConnection) Connect() (*sql.DB, error) {
	dbConnection, err := sql.Open("postgres", "user = praveen dbname = test_twichblade sslmode = disable")
	return dbConnection, err
}
