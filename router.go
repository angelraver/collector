package main

import (
	"collector/config"
	"collector/routes"
	"encoding/json"
	"net/http"
)

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data interface{} = GetResponse(r)
	if data != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", config.Get("HOST"))
		json.NewEncoder(w).Encode(data)
	} else {
		http.Error(w, "No sé de qué me estás hablando.", http.StatusNotFound)
	}
}

func GetResponse(r *http.Request) interface{} {
	switch r.Method {
	case "GET":
		return routes.GET(r)
	case "POST":
		return routes.POST(r)
	case "PUT":
		return routes.PUT(r)
	case "DELETE":
		return routes.DELETE(r)
	default:
		return nil
	}
}
