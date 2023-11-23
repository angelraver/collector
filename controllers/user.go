package controllers

import (
	"coleccionista/config"
	"coleccionista/entities"
	"coleccionista/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func UserLogin(r *http.Request, w http.ResponseWriter) string {
	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)
	results := models.UserLogin(data["user"].(string), data["password"].(string))
	defer results.Close()

	var user = entities.User{Id: 0, Name: ""}
	if results.Next() {
		err := results.Scan(&user.Id, &user.Name)
		if err != nil {
			fmt.Println(err)
		}
	}

	if user.Id == 0 {
		return "Loggin fail"
	} else {
		config.SetCookie(w, "session_key", "pepe", 3600) // Expires in 1 hour
		return "Loggin SUCCESS"
	}
}

func UserCreate(name string, password string) string {
	return models.UserCreate(name, password)
}

func UserLogout(r *http.Request, w http.ResponseWriter) string {
	config.SetCookie(w, "session_key", "pepe", -1) // Expires in 1 hour
	expiredCookie := &http.Cookie{
		Name:    "session_key",
		MaxAge:  -1,
		Expires: time.Now().Add(-time.Hour),
	}
	http.SetCookie(w, expiredCookie)
	return "Logged out"
}
