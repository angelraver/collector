package routes

import (
	"coleccionista/controllers"
	"coleccionista/entities"
	"coleccionista/shared"

	"encoding/json"
	"net/http"
	"strings"
)

func DELETE(r *http.Request, w http.ResponseWriter, authorized bool) interface{} {
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
		return controllers.ItemDelete(item)
	case "collection":
		var collection entities.Collection
		json.NewDecoder(r.Body).Decode(&collection)
		return controllers.CollectionDelete(collection)
	default:
		return nil
	}
}
