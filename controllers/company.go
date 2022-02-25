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

func CompanysUpdate(company models.Company) string {
	return dataBase.CompanyUpdate(company)
}
