package controllers

import (
	"collector/dataBase"
	"collector/models"
	"database/sql"
	"strconv"
)

// just receives a company struct and pass it to the db layer
func CompanyAdd(company models.Company) string {
	return dataBase.CompanyAdd(company)
}

func CompanyGet(id string) []models.Company {
	idInt, error := strconv.Atoi(id)
	if error != nil {
		idInt = 0
	}
	var results *sql.Rows = dataBase.CompanyGet(idInt)
	var companys []models.Company
	for results.Next() {
		var company models.Company
		err := results.Scan(&company.Id, &company.Title)
		if err != nil {
			return nil
		}
		companys = append(companys, company)
	}
	return companys
}

func CompanyUpdate(company models.Company) string {
	return dataBase.CompanyUpdate(company)
}

// receives id string and returns ok / ko
func CompanyDelete(id string) string {
	idInt, error := strconv.Atoi(id)
	if error != nil {
		return "ko"
	}
	return dataBase.CompanyDelete(idInt)
}
