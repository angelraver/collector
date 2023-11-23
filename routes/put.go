package routes

import (
	"coleccionista/controllers"
	"coleccionista/entities"
	"coleccionista/shared"

	"encoding/json"
	"net/http"
	"strings"
)

func PUT(r *http.Request, w http.ResponseWriter, authorized bool) interface{} {
	if !authorized {
		return shared.UnauthorizedMessage
	}

	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		return "Invalid URL"
	}
	entity := parts[1]
	switch entity {
	case "item":
		var item entities.Item
		json.NewDecoder(r.Body).Decode(&item)
		return controllers.ItemUpdate(item)
	case "itemtype":
		var itemtype entities.ItemType
		json.NewDecoder(r.Body).Decode(&itemtype)
		return controllers.ItemTypeUpdate(itemtype)
	default:
		return nil
	}
}
