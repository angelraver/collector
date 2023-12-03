package models

import (
	"coleccionista/dataBase"
	"database/sql"
)

func ItemTypeGet(idUser *int, id *int) *sql.Rows {
	var db *sql.DB = dataBase.Conectar()
	results, err := db.Query("SELECT * FROM itemtypesget($1, $2)", idUser, id)
	if err != nil {
		return nil
	}
	// defer db.Close()
	return results
}

func ItemTypeCreate(idUser int, name string) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL itemtypesinsert($1, $2)", idUser, name)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return name + " saved."
}

func ItemTypeUpdate(id int, name string) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL itemtypesupdate($1, $2)", id, name)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return name + " udated."
}
