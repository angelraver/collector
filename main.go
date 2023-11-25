package main

import (
	"fmt"
	"net/http"
)

func main() {
	var router Router
	fmt.Println("Starting server on 8001...")
	http.ListenAndServe(":8001", router)
}
