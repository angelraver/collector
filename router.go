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
	var p string = r.URL.Path
	switch r.Method {
	case "GET":
		return routes.GET(p, r)
	case "POST":
		return routes.POST(p, r)
	case "PUT":
		return routes.PUT(p, r)
	default:
		return nil
	}
}
