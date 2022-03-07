package routes

import (
	"collector/controllers"
	"collector/models"
	"encoding/json"
	"net/http"
)

func PUT(r *http.Request) interface{} {
	switch r.URL.Path {
	case "/gameupdate":
		var game models.Game
		json.NewDecoder(r.Body).Decode(&game)
		return controllers.GameUpdate(game)
	case "/companyupdate":
		var company models.Company
		json.NewDecoder(r.Body).Decode(&company)
		return controllers.CompanyUpdate(company)
	default:
		return nil
	}
}
