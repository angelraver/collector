package routes

import (
	"collector/config"
	"collector/controllers"
	"collector/models"
	"encoding/json"
	"net/http"
)

func POST(path string, r *http.Request) interface{} {
	switch path {
	case "/gameadd":
		var game models.Game
		json.NewDecoder(r.Body).Decode(&game)
		return controllers.GameAdd(game)
	case "/imageadd":
		return controllers.ImageAdd(r, config.Get("UPLOAD_FOLDER"))
	default:
		return nil
	}
}
