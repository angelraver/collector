package dataBase

import (
	"collector/models"
	"database/sql"
	"fmt"
)

func CompanyAdd(company models.Company) string {
	var db *sql.DB = Conectar()
	insert, err := db.Query("CALL companyAdd(?)", company.Title)
	if err != nil {
		fmt.Println(err)
		return "ko"
	}
	defer insert.Close()
	return "ok"
}

func CompanyGet(id int) *sql.Rows {
	var db *sql.DB = Conectar()
	results, err := db.Query("CALL companyGet(?)", id)
	defer db.Close()
	if err != nil {
		return nil
	}
	return results
}

func CompanyUpdate(company models.Company) string {
	var db *sql.DB = Conectar()
	rows, err := db.Query("CALL companyUpdate(?,?)",
		company.Id,
		company.Title)
	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return "ok"
}

func CompanyDelete(id int) string {
	var db *sql.DB = Conectar()
	delete, err := db.Query("CALL companyDelete(?)", id)
	if err != nil {
		fmt.Println(err)
		return "ko"
	}
	defer delete.Close()
	return "ok"
}
