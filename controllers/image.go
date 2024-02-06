package controllers

import (
	"coleccionista/entities"
	"coleccionista/models"
	"coleccionista/shared"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)
const uploadDir = "./uploads/"

func ImageServe(r *http.Request, w http.ResponseWriter, fileName string) {
	filePath := path.Join(uploadDir, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return 
	}

	http.ServeFile(w, r, filePath)
}

func ImageUpload(r *http.Request, w http.ResponseWriter) string {
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
	idItem := r.FormValue("iditem")
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

	ImageCreate(fileName, shared.IntOrNil(idUser), shared.IntOrNil(idItem))

	return fileName
}

func generateFileName(userID string) string {
	currentTime := time.Now()
	return fmt.Sprintf("%s%d%02d%02d%02d%02d%02d.jpg",
		userID,
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
	)
}

func ImageGet(idUser *int, idItem *int) []entities.Image {
	var results *sql.Rows = models.ImageGet(idUser, idItem)
	var images []entities.Image
	for results.Next() {
		var image entities.Image
		err := results.Scan(
			&image.Id,
			&image.IdItem,
			&image.Name,
			&image.CreatedAt,
			&image.UpdatedAt,
		)
		if err != nil {
			return nil
		}
		images = append(images, image)
	}
	return images
}

func ImageCreate(name string, idUser *int, idItem *int) string {
	if idUser != nil && idItem != nil {
		return models.ImageCreate(name, idUser, idItem)
	}
	return "ko"
}

func ImageDelete(r *http.Request, w http.ResponseWriter, image entities.Image) string {
	filePath := path.Join(uploadDir, image.Name)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// http.NotFound(w, r)
		return "error 1"
	}

	err := os.Remove(filePath)
	if err != nil {
		// http.Error(w, "Unable to remove the file", http.StatusInternalServerError)
		return "error 2"
	}

	return models.ImageDelete(image.Id, image.IdUser)
}