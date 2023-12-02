package dataBase

import (
	"coleccionista/config"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
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
	dbPool, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	return dbPool
}

func ConnectTCPSocket() *sql.DB {
	var (
		dbUser    = config.Get("POSTGRE_USER")
		dbPwd     = config.Get("POSTGRE_PASS")
		dbTCPHost = config.Get("POSTGRE_SERVER")
		dbPort    = config.Get("POSTGRE_PORT")
		dbName    = config.Get("POSTGRE_DBNAME")
	)

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", dbTCPHost, dbUser, dbPwd, dbPort, dbName)
	dbPool, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		fmt.Println("errorrrrr!!!!!")
		fmt.Println(err)
		panic(err.Error())
	}

	return dbPool
}
