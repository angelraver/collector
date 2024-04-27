package models

import (
	"coleccionista/dataBase"
	"database/sql"
)

func UserLogin(name string, password string) *sql.Rows {
	var db *sql.DB = dataBase.Conectar()
	results, err := db.Query("SELECT * FROM userlogin($1, $2)", name, password)
	if err != nil {
		return nil
	}
	defer db.Close()
	return results
}

func UserCreate(name string, password string) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL usersinsert($1, $2)", name, password)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return name + " saved."
}

func UserGetByName(name string) *sql.Rows {
	var db *sql.DB = dataBase.Conectar()
	results, err := db.Query("SELECT * FROM usergetbyname($1)", name)
	if err != nil {
		return nil
	}
	defer db.Close()
	return results
}
