package main

import (
	"encoding/json"
	"log"
	"os"
)

var (
	applicationSettings appSettings
)

type appSettings struct {
	Token                  string `json:"token"`
	AppId                  string `json:"appId"`
	SupergoonGamesServerId string `json:"supergoonGamesServerId"`
	GithubAccessToken      string `json:"github_access_token"`
}

func initializeAppSettings() {

	jsonContent, err := os.ReadFile("appsettings.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
		return
	}
	err = json.Unmarshal(jsonContent, &applicationSettings)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
		return
	}
}
