package controllers

import (
	"coleccionista/config"
	"coleccionista/entities"
	"coleccionista/models"
	"encoding/json"
	"net/http"
	"time"
)

func UserLogin(r *http.Request, w http.ResponseWriter) map[string]interface{} {
	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)
	results := models.UserLogin(data["user"].(string), data["password"].(string))
	defer results.Close()

	var user = entities.User{Id: 0, Name: ""}
	if results.Next() {
		err := results.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil
		}
	}

	if user.Id == 0 {
		response := map[string]interface{}{
			"message": "Login FAIL",
		}
		return response
	} else {
		// config.SetCookie(w, "iduser", strconv.Itoa(user.Id), 3600) // Expires in 1 hour
		response := map[string]interface{}{
			"message": "Login SUCCESS",
			"iduser":  user.Id,
		}

		// sessionKey, err := r.Cookie("iduser")
		// fmt.Println("sessionKey: ")
		// fmt.Println(sessionKey)

		// if err != nil {
		// 	fmt.Println("ERROR EN COOKIE:")
		// }

		return response
	}
}

func UserGetByName(userName string) (entities.User, error) {
	var user = entities.User{Id: 0, Name: "", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	results := models.UserGetByName(userName)
	if results.Next() {
		err := results.Scan(&user.Id, &user.Name, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return user, nil
		}
	}

	return user, nil
}

func UserCreate(name string, password string) string {
	return models.UserCreate(name, password)
}

func UserLogout(r *http.Request, w http.ResponseWriter) string {
	config.SetCookie(w, "iduser", "", -1) // Expires in 1 hour
	expiredCookie := &http.Cookie{
		Name:    "iduser",
		MaxAge:  -1,
		Expires: time.Now().Add(-time.Hour),
	}
	http.SetCookie(w, expiredCookie)
	return "Logged out"
}
