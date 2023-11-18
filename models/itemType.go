package models

import (
	"coleccionista/dataBase"
	"database/sql"
	"fmt"
)

func ItemTypeGet(idUser *int, id *int) *sql.Rows {
	var db *sql.DB = dataBase.Conectar()
	results, err := db.Query("SELECT * FROM public.itemtypesget($1, $2)", idUser, id)
	if err != nil {
		fmt.Println("ERROR IN MODEL")
		return nil
	}
	// defer db.Close()
	fmt.Println("RESULTS IN MODEL:")
	fmt.Println(results)
	return results
}

func ItemTypeCreate(idUser int, name string) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL public.itemtypesinsert($1, $2)", idUser, name)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return name + " saved."
}

func ItemTypeUpdate(id int, name string) string {
	var db *sql.DB = dataBase.Conectar()
	rows, err := db.Query("CALL public.itemtypesupdate($1, $2)", id, name)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return name + " udated."
}
