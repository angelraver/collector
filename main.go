package main

import (
	"coleccionista/routes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data interface{} = GetResponse(r, w)
	if data != nil {
		w.Header().Set("Content-Type", "application/json")
		// Allow requests from https://angelraver.github.io
		w.Header().Set("Access-Control-Allow-Origin", "https://angelraver.github.io")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "No sé de qué me estás hablando.", http.StatusNotFound)
	}
}
func Authorized(w http.ResponseWriter, r *http.Request) bool {
	return true
	// sessionKey, err := r.Cookie("iduser")
	// fmt.Println("sessionKey: ")
	// fmt.Println(sessionKey)
	// if err != nil {
	// 	return false
	// }

	// if sessionKey != nil {
	// 	return true
	// }

	// return false
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
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ServeHTTP)

	// Use the cors package to handle CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"https://angelraver.github.io"},
	}).Handler(mux)

	fmt.Println("Starting server on 8080...")
	http.ListenAndServe(":8080", corsHandler)
}
