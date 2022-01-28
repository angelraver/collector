package handlers

import (
	"collector/controllers"
	"collector/models"
	"encoding/json"
	"net/http"
)

func GameAdd(w http.ResponseWriter, r *http.Request) {
	var game models.Game
	json.NewDecoder(r.Body).Decode(&game)
	result := controllers.GameAdd(game)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GameGet(w http.ResponseWriter, r *http.Request) {
	var id string = r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controllers.GameGet(id))
}

func GameUpdate(w http.ResponseWriter, r *http.Request) {
	var game models.Game
	json.NewDecoder(r.Body).Decode(&game)
	result := controllers.GameUpdate(game)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
