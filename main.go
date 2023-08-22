package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/kjblanchard/sggDiscordBot/github"
)

type appSettings struct {
	Token                  string `json:"token"`
	AppId                  string `json:"appId"`
	SupergoonGamesServerId string `json:"supergoonGamesServerId"`
	GithubAccessToken      string `json:"github_access_token"`
}

func main() {
	jsonContent, err := os.ReadFile("appsettings.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
		return
	}
	var config appSettings
	err = json.Unmarshal(jsonContent, &config)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
		return
	}
	s, _ := discordgo.New("Bot " + config.Token)
	github.InitializeGithub(config.GithubAccessToken)
	_, err = s.ApplicationCommandBulkOverwrite(config.AppId, config.SupergoonGamesServerId, []*discordgo.ApplicationCommand{
		{
			Name:        "check-repos",
			Description: "Checks to see all of the repos available",
		},
		{
			Name:        "hello-world",
			Description: "Showcase of a basic slash command",
		},
	})
	if err != nil {
		log.Fatal("Error doing the thing", err)
		// Handle the error
	}
	s.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		data := i.ApplicationCommandData()
		switch data.Name {
		case "hello-world":
			err := s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Hello world!",
					},
				},
			)
			if err != nil {
				log.Fatal("Error registering", err)
				// Handle the error
			}
		case "check-repos":
			roles := i.Member.Roles

			github.GetAllRepos()

			for _, role := range roles {
				log.Print(role)
			}
			err := s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Heres a list of all of Kevins Repos..",
						Embeds: []*discordgo.MessageEmbed{
							{
								Title: "Repo",
							},

						},
					},
				},
			)
			if err != nil {
				log.Fatal("Error registering", err)
				// Handle the error
			}
		}
	})
	err = s.Open()
	if err != nil {
		log.Fatal("Error registering", err)
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	err = s.Close()
	if err != nil {
		log.Fatal("Error registering", err)
	}
}
