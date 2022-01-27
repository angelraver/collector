package main

import (
	"collector/handlers"
	"net/http"
)

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handlers.Home(w, r)
	case "/gameget":
		handlers.GamesGet(w, r)
	// case "/userform":
	// 	handlers.UserForm(w, r)
	// case "/userjson":
	// 	handlers.UserJson(w, r)
	// case "/getusers":
	// 	handlers.GetUsers(w, r, Conectar())
	default:
		http.Error(w, "Woops! nothing to see here...", http.StatusNotFound)
	}
}
