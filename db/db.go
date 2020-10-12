package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToDB(host, user, password, database string, port int) (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)
	db, err := sql.Open("postgres", dbinfo)
	return db, err
}
