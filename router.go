package main

import (
	"coleccionista/config"
	"coleccionista/routes"
	"encoding/json"
	"net/http"
)

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data interface{} = GetResponse(r, w)
	if data != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", config.Get("HOST"))
		json.NewEncoder(w).Encode(data)
	} else {
		http.Error(w, "No sé de qué me estás hablando.", http.StatusNotFound)
	}
}

func Authorized(w http.ResponseWriter, r *http.Request) bool {
	sessionKey, err := r.Cookie("session_key")
	if err != nil {
		return false
	}

	if sessionKey != nil {
		return true
	}

	return false
}

func GetResponse(r *http.Request, w http.ResponseWriter) interface{} {
	var authorized bool = Authorized(w, r)
	switch r.Method {
	case "GET":
		return routes.GET(r, w, authorized)
	case "POST":
		return routes.POST(r, w, authorized)
	case "PUT":
		return routes.PUT(r, w, authorized)
	default:
		return nil
	}
}
