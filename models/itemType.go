package models

import (
	"coleccionista/dataBase"
	"database/sql"
)

func ItemTypeGet(idUser *int, id *int) *sql.Rows {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("SELECT * FROM itemtypesget($1, $2)", idUser, id)
	if err != nil {
		return nil
	}
	defer db.Close()
	return rows
}

func ItemTypeCreate(idUser int, name string) *sql.Row {
	var db *sql.DB = dataBase.Conectar()
	defer db.Close()
	return db.QueryRow("SELECT * FROM itemtypesinsert($1, $2) AS id", idUser, name)
}

func ItemTypeUpdate(id int, name string) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL itemtypesupdate($1, $2)", id, name)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return name + " updated."
}

func ItemTypeDelete(id int, idUser int) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL itemtypesdelete($1, $2)", id, idUser)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return "collection deleted."
}