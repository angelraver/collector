package dataBase

import (
	"coleccionista/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Returns an open connection
func Conectar() *sql.DB {
	var (
		host     = config.Get("POSTGRE_SERVER")
		port     = config.Get("POSTGRE_PORT")
		user     = config.Get("POSTGRE_USER")
		password = config.Get("POSTGRE_PASS")
		dbname   = config.Get("POSTGRE_DBNAME")
	)
	psqlInfo := fmt.Sprintf("host=%s port="+port+" user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	return db
}