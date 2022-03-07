package routes

import (
	"collector/controllers"
	"net/http"
)

func DELETE(r *http.Request) interface{} {
	switch r.URL.Path {
	case "/companydelete":
		var id string = r.URL.Query().Get("id")
		return controllers.CompanyDelete(id)
	default:
		return nil
	}
}
