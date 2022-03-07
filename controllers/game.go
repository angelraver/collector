package controllers

import (
	"collector/dataBase"
	"collector/models"
	"database/sql"
)

// Pass a game stuct to the database layer
func GameAdd(game models.Game) string {
	return dataBase.GameAdd(game)
}

func GameGet(id string) []models.Game {
	var results *sql.Rows = dataBase.GameGet(id)
	var games []models.Game
	for results.Next() {
		var game models.Game
		err := results.Scan(&game.Id,
			&game.IdConsole,
			&game.Title,
			&game.Stars,
			&game.Qty)
		if err != nil {
			return nil
		}
		games = append(games, game)
	}
	return games
}

func GameUpdate(game models.Game) string {
	return dataBase.GameUpdate(game)
}

// func GameDelete(game models.Game) string {
// 	return dataBase.GameDelete(game)
// }
