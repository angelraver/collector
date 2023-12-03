package models

import (
	"coleccionista/dataBase"
	"database/sql"
)

func ItemGet(idUser *int, id *int, idItemType *int) *sql.Rows {
	var db *sql.DB = dataBase.Conectar()
	results, err := db.Query("SELECT * FROM itemsget($1, $2, $3)", idUser, id, idItemType)
	if err != nil {
		return nil
	}
	defer db.Close()
	return results
}

func ItemCreate(idUser int, idItemType int, title string) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL itemsinsert($1, $2, $3)", idUser, idItemType, title)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return title + " saved."
}

func ItemUpdate(idItem int, idItemType int, title string) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL itemsupdate($1, $2, $3)", idItem, idItemType, title)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return title + " udated."
}
