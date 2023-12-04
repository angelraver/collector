package main

import (
	"net/http"

	// "github.com/rs/cors"
	"github.com/gorilla/handlers"
)

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", ServeHTTP)
	// corsHandler := cors.Default().Handler(mux)
	// fmt.Println("Starting server on 8080...")
	// http.ListenAndServe(":8080", corsHandler)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
