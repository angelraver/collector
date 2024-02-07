package controllers

import (
	"coleccionista/entities"
	"coleccionista/models"
	"coleccionista/shared"
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
)

const BucketName = "coleccionista-bucket"

func ImageUpload(r *http.Request, w http.ResponseWriter) string {
	idUser := r.FormValue("iduser")
	idItem := r.FormValue("iditem")
	fileName := generateFileName(idUser)

	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Print(handler)
		fmt.Fprintf(w, "Error retrieving uploaded file: %v", err) // Use w for response
		return "error 1"
	}
	defer file.Close()

	// Connect to Google Cloud Storage
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
			fmt.Fprintf(w, "Error creating storage client: %v", err)
			return "error 2"
	}

	bucket := client.Bucket(BucketName)

	// Create a new object in the bucket
	obj := bucket.Object(fileName)
	ww := obj.NewWriter(ctx)

	// Copy the file contents to the object
	if _, err := io.Copy(ww, file); err != nil {
			fmt.Fprintf(w, "Error uploading file to GCS: %v", err)
			return "error 3"
	}

	// Close the writer to finalize the upload
	if err := ww.Close(); err != nil {
			fmt.Fprintf(ww, "Error closing writer: %v", err)
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
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		// fmt.Errorf("error creating storage client: %v", err)
		return "error 1"
	}

	bucket := client.Bucket(BucketName)
	obj := bucket.Object(image.Name)

	if err := obj.Delete(ctx); err != nil {
		// fmt.Errorf("error deleting object: %v", err)
		// return "error 2"
	}

	return models.ImageDelete(image.Id, image.IdUser)
}
