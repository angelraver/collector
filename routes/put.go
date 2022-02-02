package routes

import (
	"collector/controllers"
	"collector/models"
	"encoding/json"
	"net/http"
)

func PUT(path string, r *http.Request) interface{} {
	switch path {
	case "/gameupdate":
		var game models.Game
		json.NewDecoder(r.Body).Decode(&game)
		return controllers.GameUpdate(game)
	default:
		return nil
	}
}
