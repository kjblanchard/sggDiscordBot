package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	S                      *discordgo.Session
	discordApplicationId   string
	supergoonGamesServerId string
)

func InitializeDiscord(token string, appId string, supergoonServerId string) {
	S, _ = discordgo.New("Bot " + token)
	discordApplicationId = appId
	supergoonGamesServerId = supergoonServerId
}

func OpenDiscordWebsocketConnection() {
	err := S.Open()
	if err != nil {
		log.Fatal("Error Opening websocket connection for discord bot!\nError: ", err)
	}
}

func CloseDiscord() {
	err := S.Close()
	if err != nil {
		log.Fatal("Error Closing bot", err)
	}
}
