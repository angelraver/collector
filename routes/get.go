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
	case "collection":
		if !authorized {
			return shared.UnauthorizedMessage
		}
		return controllers.CollectionGet(shared.IntOrNil(param1), shared.IntOrNil(param2))
	case "collectionbkp":
		if !authorized {
			return shared.UnauthorizedMessage
		}
		return controllers.CollectionBkp()
	case "games":
		if !authorized {
			return shared.UnauthorizedMessage
		}
		return controllers.IgdbGetGames(param1, param2)
	case "item":
		if !authorized {
			return shared.UnauthorizedMessage
		}
		return controllers.ItemGet(shared.IntOrNil(param1), shared.IntOrNil(param2), shared.IntOrNil(param3))
	case "itembkp":
		if !authorized {
			return shared.UnauthorizedMessage
		}
		return controllers.ItemBkp()
	case "itemimage":
		if !authorized {
			return shared.UnauthorizedMessage
		}
		return controllers.ImageGet(shared.IntOrNil(param1), shared.IntOrNil(param2))
	case "logout":
		return controllers.UserLogout(r, w)
	default:
		return nil
	}
}
