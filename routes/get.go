package routes

import (
	"coleccionista/controllers"
	"coleccionista/shared"
	"net/http"
	"strings"
)

type HomeData struct {
	CurrentDate string
}

func GET(r *http.Request, w http.ResponseWriter, authorized bool) interface{} {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		return "Invalid URL"
	}
	entity := parts[1]
	param1 := ""
	param2 := ""
	param3 := ""
	if len(parts) >= 3 {
		param1 = parts[2]
	}
	if len(parts) >= 4 {
		param2 = parts[3]
	}
	if len(parts) >= 5 {
		param3 = parts[4]
	}

	switch entity {
	case "item":
		if !authorized {
			return shared.UnauthorizedMessage
		}
		return controllers.ItemGet(intOrNil(param1), intOrNil(param2), intOrNil(param3))
	case "itemtype":
		if !authorized {
			return shared.UnauthorizedMessage
		}
		return controllers.ItemTypeGet(intOrNil(param1), intOrNil(param2))
	case "game":
		if !authorized {
			return shared.UnauthorizedMessage
		}
		return controllers.GameGet(param1, intOrNil(param2))
	case "logout":
		return controllers.UserLogout(r, w)
	default:
		return nil
	}
}
