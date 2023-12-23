package controllers

import (
	"coleccionista/config"
	"coleccionista/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

func IgdbRequest(entity string, param string, param2 string) interface{} {
	baseURL := "https://api.igdb.com/v4/"
	bodyParams := ""

	switch entity {
	case "games":
		bodyParams = "fields id, name, first_release_date; search \"" + param +"\";"
		if len(param2) > 0 && param2 != "0" {
			bodyParams = bodyParams + "where platforms = [" + param2 +"];"
		}
		bodyParams = bodyParams + "limit 20;"
	case "covers":
		bodyParams = "fields url; where game = " + param + ";"
	default:
		return nil
	}

	request, err := http.NewRequest("POST", baseURL + entity, strings.NewReader(bodyParams))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	request.Header.Set("Authorization", "Bearer " + config.Get("CLIENT_TOKEN"))
	request.Header.Set("Client-ID", config.Get("CLIENT_ID"))
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	body, err := ioutil.ReadAll(response.Body)
	return body
}

func IgdbGetGames(name string, idPlatform string) []entities.IGDBGameResponse {
	bodyInterface := IgdbRequest("games", name, idPlatform)
	body, ok := bodyInterface.([]byte)
	if !ok {
		fmt.Println("Error: Could not convert response body to []byte")
		return nil
	}
	var gameResponses []entities.IGDBGameResponse
	json.Unmarshal(body, &gameResponses)
	return gameResponses
}

func IgdbGetCover(id string) string {
	bodyInterface := IgdbRequest("covers", id, "")
	body, ok := bodyInterface.([]byte)
	if !ok {
		fmt.Println("Error: Could not convert response body to []byte")
	}
	var coverResponses []entities.IGDBCoverResponse
	json.Unmarshal(body, &coverResponses)
	coverUrl := coverResponses[0].Url
	return extractFilenameFromURL("http:" + coverUrl)
}

func extractFilenameFromURL(inputURL string) string {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return ""
	}

	filename := filepath.Base(parsedURL.Path)
	return filename
}

