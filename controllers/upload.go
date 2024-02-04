package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

func ImageUpload(w http.ResponseWriter, r *http.Request) string {
	const uploadDir = "./uploads/"
	// Ensure the uploads directory exists
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, 0755)
	}

	err := r.ParseMultipartForm(10 << 20) // Set a limit for the uploaded file size

	if err != nil {
		// http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return "error 1"
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		// http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		fmt.Print(handler)
		return "error 2"
	}

	defer file.Close()

	idUser := r.FormValue("iduser")
	fileName := generateFileName(idUser)

	// Create the file in the uploads directory
	filePath := path.Join(uploadDir, fileName)
	f, err := os.Create(filePath)
	if err != nil {
		// http.Error(w, "Unable to save the file", http.StatusInternalServerError)
		return "error 3"
	}
	defer f.Close()

	// Copy the file content to the created file
	_, err = io.Copy(f, file)
	if err != nil {
		// http.Error(w, "Unable to save the file", http.StatusInternalServerError)
		return "error 4"
	}

	return fileName
}

func generateFileName(userID string) string {
	currentTime := time.Now()
	return fmt.Sprintf("%s-%d-%02d-%02d_%02d-%02d-%02d.jpg",
		userID,
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
	)
}
