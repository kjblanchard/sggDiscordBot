package main

import (
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
	applicationSettings.Token = os.Getenv("DISCORD_BOT_TOKEN")
	applicationSettings.AppId = os.Getenv("DISCORD_APP_ID")
	applicationSettings.SupergoonGamesServerId = os.Getenv("DISCORD_SUPERGOON_GUILD_ID")
	applicationSettings.GithubAccessToken = os.Getenv("GITHUB_ACCESS_TOKEN")

}
