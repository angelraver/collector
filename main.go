package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	var router Router
	fmt.Println("Starting server on 8001...")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.ListenAndServe(":8001", handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
