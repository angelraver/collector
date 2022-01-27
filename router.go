package main

import (
	"collector/handlers"
	"net/http"
)

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err = false
	switch r.Method {
	case "GET":
		switch r.URL.Path {
		case "/":
			handlers.Home(w, r)
		case "/gameget":
			handlers.GameGet(w, r)
		default:
			err = true
		}
	case "POST":
		switch r.URL.Path {
		case "/gameadd":
			handlers.GameAdd(w, r)
		default:
			err = true
		}
	}
	if err {
		http.Error(w, "404 No sé de qué me estás hablando.", http.StatusNotFound)
		return
	}
}
