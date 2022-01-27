package dataBase

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Returns an open connection
func Conectar() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/colector")
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()
	fmt.Println("connected to db!")
	return db
}
