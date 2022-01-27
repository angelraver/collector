package handlers

import (
	"collector/dataBase"
	"collector/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func GameAdd(w http.ResponseWriter, r *http.Request) {
	var db *sql.DB = dataBase.Conectar()
	var game models.Game
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&game)
	insert, err := db.Query("CALL gameAdd(?,?,?,?)",
		game.Title,
		game.IdConsole,
		game.Stars,
		game.Qty)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	if insert != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("ok")
	}
}

func GameGet(w http.ResponseWriter, r *http.Request) {
	var db *sql.DB = dataBase.Conectar()
	results, err := db.Query("CALL gameGet(?)", nil)
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}

	var games []models.Game
	for results.Next() {
		var game models.Game
		// for each row, scan the result into our tag composite object
		err := results.Scan(&game.Id, &game.IdConsole, &game.Title, &game.Stars, &game.Qty)
		if err != nil {
			panic(err.Error())
		}
		games = append(games, game)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}
