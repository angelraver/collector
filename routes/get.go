package routes

import (
	"coleccionista/controllers"
	"coleccionista/shared"
	"net/http"
	"strings"
)

func GET(r *http.Request, w http.ResponseWriter, authorized bool) interface{} {
	if !authorized {
		return shared.UnauthorizedMessage
	}

	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		return "Invalid URL"
	}
	entity := parts[1]
	param1 := ""
	param2 := ""
	if len(parts) >= 3 {
		param1 = parts[2]
	}
	if len(parts) >= 4 {
		param2 = parts[3]
	}

	switch entity {
	case "item":
		return controllers.ItemGet(intOrNil(param1), intOrNil(param2))
	case "itemtype":
		return controllers.ItemTypeGet(intOrNil(param1), intOrNil(param2))
	case "logout":
		return controllers.UserLogout(r, w)
	default:
		return nil
	}
}
