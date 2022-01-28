package main

import (
	"collector/handlers"
	"net/http"
)

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var path string = r.URL.Path
	switch r.Method {
	case "GET":
		switch path {
		case "/":
			handlers.Home(w, r)
		case "/gameget":
			handlers.GameGet(w, r)
		default:
			Bad(w, r)
		}
	case "POST":
		switch path {
		case "/gameadd":
			handlers.GameAdd(w, r)
		default:
			Bad(w, r)
		}
	case "PUT":
		switch path {
		case "/gameupdate":
			handlers.GameUpdate(w, r)
		default:
			Bad(w, r)
		}
	default:
		Bad(w, r)
	}
}

func Bad(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "No sé de qué me estás hablando.", http.StatusNotFound)
}
