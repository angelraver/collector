package controllers

import (
	"coleccionista/config"
	"coleccionista/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GameGet(name string, idPlatform *int) []entities.GameResponse {
	baseURL := "https://api.igdb.com/v4/games"
	bodyParams := "fields id, name; search \"" + name +"\";"
	if (idPlatform != nil) {
		bodyParams = bodyParams + "where (platforms = [" + strconv.Itoa(*idPlatform) +"]);"
	}
	bodyParams = bodyParams + "limit 20;"
	request, err := http.NewRequest("POST", baseURL, strings.NewReader(bodyParams))
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

	var gameResponses []entities.GameResponse
	json.Unmarshal(body, &gameResponses)
	return gameResponses
}

