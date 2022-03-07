package routes

import (
	"collector/controllers"
	"net/http"
)

func GET(r *http.Request) interface{} {
	switch r.URL.Path {
	case "/gameget":
		var id string = r.URL.Query().Get("id")
		return controllers.GameGet(id)
	case "/companyget":
		var id string = r.URL.Query().Get("id")
		return controllers.CompanyGet(id)
	default:
		return nil
	}
}
