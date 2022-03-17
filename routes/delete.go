package routes

import (
	"collector/controllers"
	"net/http"
)

// receives the request and returns ok / ko
func DELETE(r *http.Request) interface{} {
	switch r.URL.Path {
	case "/companydelete":
		var id string = r.URL.Query().Get("id")
		return controllers.CompanyDelete(id)
	default:
		return nil
	}
}
