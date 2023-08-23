package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	s                      *discordgo.Session
	discordApplicationId   string
	supergoonGamesServerId string
)

func InitializeDiscord(token string, appId string, supergoonServerId string) {
	s, _ = discordgo.New("Bot " + token)
	discordApplicationId = appId
	supergoonGamesServerId = supergoonServerId
}

func OpenDiscordWebsocketConnection() {
	err := s.Open()
	if err != nil {
		log.Fatal("Error Opening websocket connection for discord bot!\nError: ", err)
	}
}

func CloseDiscord() {
	err := s.Close()
	if err != nil {
		log.Fatal("Error Closing bot", err)
	}
}
