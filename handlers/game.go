package handlers

import (
	"collector/dataBase"
	"collector/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func GameJson(w http.ResponseWriter, r *http.Request) {
	// var game models.Game
	// // body, err := ioutil.ReadAll(r.Body)
	// // err = json.Unmarshal(body, &user)
	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&game)

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(game.Title)
	fmt.Println("working...")
}

// func GamesGet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
func GamesGet(w http.ResponseWriter, r *http.Request) {
	var db *sql.DB = dataBase.Conectar()
	results, err := db.Query("CALL gameGet(?)", nil)
	// defer results.Close()
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

// perform a db.Query insert
// insert, err := db.Query("CALL userCreate(?,?)", "angel", "123456")
// if err != nil {
//   panic(err.Error())
// }
// be careful deferring Queries if you are using transactions
// defer insert.Close()
