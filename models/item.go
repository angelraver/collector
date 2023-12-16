package models

import (
	"coleccionista/dataBase"
	"database/sql"
)

func ItemGet(idUser *int, id *int, idIdCollection *int) *sql.Rows {
	var db *sql.DB = dataBase.Conectar()
		results, err := db.Query("SELECT * FROM itemsget($1, $2, $3)", idUser, id, idIdCollection)
	if err != nil {
		return nil
	}
	defer db.Close()
	return results
}

func ItemCreate(
	idUser int,
	idItemType int,
	idCollection int,
	title string,
	idIgdb int,
	cover string,
	) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL itemsinsert($1, $2, $3, $4, $5, $6)", idUser, idItemType, idCollection, title, idIgdb, cover)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return title + " saved."
}

func ItemUpdate(id int, idUser int, title string) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL itemsupdate($1, $2, $3)", id, idUser, title)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return title + " updated."
}

func ItemDelete(id int, idUser int) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL itemsdelete($1, $2)", id, idUser)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return "item deleted."
}