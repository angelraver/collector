package models

import (
	"coleccionista/dataBase"
	"database/sql"
)

func ImageGet(idUser *int, idItem *int) *sql.Rows {
	var db *sql.DB = dataBase.Conectar()
		results, err := db.Query("SELECT * FROM imagesget($1, $2)", idUser, idItem)
	if err != nil {
		return nil
	}
	defer db.Close()
	return results
}

func ImageCreate(
	name string,
	idUser *int,
	idItem *int,
) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL imagesinsert($1, $2, $3)", name, idUser, idItem)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return name + " saved."
}

func ImageDelete(id int, idUser int) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL imagesdelete($1, $2)", id, idUser)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return "image deleted."
}