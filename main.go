package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// var router Router
	// corsHandler := cors.Default().Handler(router)
	// fmt.Println("Starting server on 8080...")
	// http.ListenAndServe(":8080", corsHandler)

	mux := http.NewServeMux()
	mux.HandleFunc("/", ServeHTTP)
	corsHandler := cors.Default().Handler(mux)
	fmt.Println("Starting server on 8080...")
	http.ListenAndServe(":8080", corsHandler)
}
