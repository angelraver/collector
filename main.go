package main

import (
	"fmt"
	"net/http"
)

func main() {
	var router Router
	fmt.Println("Starting server on 3000...")
	http.ListenAndServe(":3000", router)
}
