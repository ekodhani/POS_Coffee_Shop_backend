package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dbuser := "root"
	dbpass := ""
	dbname := "pos_flow_coffee"
	dbhost := "localhost"

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+dbhost+")/"+dbname)

	if err != nil {
		return nil, err
	}
	return db, nil
}
