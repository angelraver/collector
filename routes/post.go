package routes

import (
	"collector/config"
	"collector/controllers"
	"collector/models"
	"encoding/json"
	"net/http"
)

func POST(r *http.Request) interface{} {
	switch r.URL.Path {
	case "/gameadd":
		var game models.Game
		json.NewDecoder(r.Body).Decode(&game)
		return controllers.GameAdd(game)
	case "/imageadd":
		return controllers.ImageAdd(r, config.Get("UPLOAD_FOLDER"))
	case "/companyadd":
		var company models.Company
		json.NewDecoder(r.Body).Decode(&company)
		return controllers.CompanyAdd(company)
	default:
		return nil
	}
}
