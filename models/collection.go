package models

import (
	"coleccionista/dataBase"
	"database/sql"
)

func CollectionGet(idUser *int, id *int) *sql.Rows {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("SELECT * FROM collectionsget($1, $2)", idUser, id)
	if err != nil {
		return nil
	}
	defer db.Close()
	return rows
}

func CollectionCreate(idUser int, idItemType int, name string, idplatform int) *sql.Row {
	var db *sql.DB = dataBase.Conectar()
	defer db.Close()
	return db.QueryRow("SELECT * FROM collectionsinsert($1, $2, $3, $4) AS id", idUser, idItemType, name, idplatform)
}

func CollectionUpdate(id int, idUser int, name string, idplatform int) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL collectionsupdate($1, $2, $3, $4)", id, idUser, name, idplatform)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return name + " updated."
}

func CollectionDelete(id int, idUser int) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL collectionsdelete($1, $2)", id, idUser)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return "collection deleted."
}