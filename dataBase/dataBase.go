package dataBase

import (
	"collector/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Returns an open connection
func Conectar() *sql.DB {
	db, err := sql.Open("mysql", config.Get("SQL_USER")+":"+config.Get("SQL_PASSWORD")+"@tcp(127.0.0.1:3306)/collector")
	if err != nil {
		panic(err.Error())
	}
	return db
}
