package main

import "os"

type AppSettings struct {
	Token             string
	AppID             string
	GuildID           string
	GithubAccessToken string
}

var ApplicationSettings AppSettings

func InitializeAppSettings() {
	ApplicationSettings.Token = os.Getenv("DISCORD_BOT_TOKEN")
	ApplicationSettings.AppID = os.Getenv("DISCORD_APP_ID")
	ApplicationSettings.GuildID = os.Getenv("DISCORD_SUPERGOON_GUILD_ID")
	ApplicationSettings.GithubAccessToken = os.Getenv("GITHUB_ACCESS_TOKEN")
}
