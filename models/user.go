package models

import (
	"coleccionista/dataBase"
	"database/sql"
)

func UserLogin(name string, password string) *sql.Rows {
	var db *sql.DB = dataBase.Conectar()
	results, err := db.Query("SELECT * FROM public.userlogin($1, $2)", name, password)
	if err != nil {
		return nil
	}
	defer db.Close()
	return results
}

func UserCreate(name string, password string) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL public.usersinsert($1, $2)", name, password)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return name + " saved."
}
