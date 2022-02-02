package routes

import (
	"collector/controllers"
	"net/http"
)

func GET(path string, r *http.Request) interface{} {
	switch path {
	case "/gameget":
		var id string = r.URL.Query().Get("id")
		return controllers.GameGet(id)
	default:
		return nil
	}
}
