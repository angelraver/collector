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
	author string,
	year int,
	) int {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL itemsinsert($1, $2, $3, $4, $5, $6, $7, $8)", idUser, idItemType, idCollection, title, idIgdb, cover, author, year)
	if err != nil {
		return 0
	}
	defer rows.Close()

	var idResult int
	db.QueryRow("SELECT * FROM itemsgetlast($1, $2)", idUser, idCollection).Scan(&idResult)

	return idResult
}

func ItemUpdate(id int, idUser int, title string, author string, year int) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL itemsupdate($1, $2, $3, $4, $5)", id, idUser, title, author, year)
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