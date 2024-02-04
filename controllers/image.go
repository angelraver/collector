package controllers

import (
	"net/http"
	"os"
	"path"
)

func ImageServe(w http.ResponseWriter, r *http.Request, fileName string) {
	const uploadDir = "./uploads/"
	filePath := path.Join(uploadDir, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return 
	}

	http.ServeFile(w, r, filePath)
}