package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	discordbot "sgg-discord-bot/discord_bot"
	goongithub "sgg-discord-bot/goon_github"
	goonhttp "sgg-discord-bot/goon_http"
)

func main() {
	InitializeAppSettings()

	session := discordbot.InitializeDiscord(ApplicationSettings.Token, ApplicationSettings.AppID, ApplicationSettings.GuildID)
	goongithub.InitializeGithub(ApplicationSettings.GithubAccessToken)
	discordbot.AddAllSlashCommands(session, ApplicationSettings.GuildID)

	err := discordbot.OpenDiscordConnection(session)
	if err != nil {
		log.Fatalf("Error opening discord connection: %v", err)
	}

	goonhttp.StartServer()

	log.Println("Press Ctrl+C to exit")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	discordbot.CloseDiscord(session)
}
