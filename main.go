package main

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
)

type appSettings struct {
	Token                  string `json:"token"`
	AppId                  string `json:"appId"`
	SupergoonGamesServerId string `json:"supergoonGamesServerId"`
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
	_, err = s.ApplicationCommandBulkOverwrite(config.AppId, config.SupergoonGamesServerId, []*discordgo.ApplicationCommand{
		{
			Name:        "hello-world",
			Description: "Showcase of a basic slash command",
		},
	})
	if err != nil {
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
				// Handle the error
			}
		}
	})
	err = s.Open()
	if err != nil {
		// Handle the error
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	err = s.Close()
	if err != nil {
		// Handle the error
	}
}
