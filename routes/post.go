package routes

import (
	"coleccionista/controllers"
	"coleccionista/entities"
	"coleccionista/shared"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func intOrNil(param string) *int {
	if len(param) > 0 {
		value, err := strconv.Atoi(param)
		if err != nil {
			fmt.Println(err)
			return nil
		} else {
			return &value
		}
	}
	return nil
}

func POST(r *http.Request, w http.ResponseWriter, authorized bool) interface{} {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		return "Invalid URL"
	}
	useCase := parts[1]

	if useCase == "login" {
		return controllers.UserLogin(r, w)
	}

	if !authorized {
		return shared.UnauthorizedMessage
	}

	switch useCase {
	case "item":
		var item entities.Item
		json.NewDecoder(r.Body).Decode(&item)
		return controllers.ItemCreate(item)
	case "itemtype":
		var itemtype entities.ItemType
		json.NewDecoder(r.Body).Decode(&itemtype)
		return controllers.ItemTypeCreate(itemtype)
	default:
		return nil
	}
}
