package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/kjblanchard/sggDiscordBot/github"
	"github.com/kjblanchard/sggDiscordBot/discord"
)


func main() {
	initializeAppSettings()
	discord.InitializeDiscord(applicationSettings.Token, applicationSettings.AppId, applicationSettings.SupergoonGamesServerId)
	github.InitializeGithub(applicationSettings.GithubAccessToken)
	discord.AddAllSlashCommands()
	discord.OpenDiscordWebsocketConnection()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
	discord.CloseDiscord()

}
