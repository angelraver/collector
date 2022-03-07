package dataBase

import (
	"collector/models"
	"database/sql"
)

func GameAdd(game models.Game) string {
	var db *sql.DB = Conectar()
	insert, err := db.Query("CALL gameAdd(?,?,?,?)",
		game.Title,
		game.IdConsole,
		game.Stars,
		game.Qty)
	if err != nil {
		return "ko"
	}
	defer insert.Close()
	return "ok"
}

func GameGet(id string) *sql.Rows {
	var db *sql.DB = Conectar()
	results, err := db.Query("CALL gameGet(?)", id)
	defer db.Close()
	if err != nil {
		return nil
	}
	return results
}

func GameUpdate(game models.Game) string {
	var db *sql.DB = Conectar()
	rows, err := db.Query("CALL gameUpdate(?,?,?,?,?)",
		game.Id,
		game.Title,
		game.IdConsole,
		game.Stars,
		game.Qty)

	if err != nil {
		return "ko"
	}
	defer rows.Close()
	return "ok"
}

func GameDelete(id string) *sql.Rows {
	var db *sql.DB = Conectar()
	results, err := db.Query("CALL gameDelete(?)", id)
	defer db.Close()
	if err != nil {
		return nil
	}
	return results
}
