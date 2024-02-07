package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ServeHTTP)
	corsHandler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}).Handler(mux)
	fmt.Println("Starting server on 8080...")
	http.ListenAndServe(":8080", corsHandler)
}
