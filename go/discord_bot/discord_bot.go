package discord_bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var BotInstance *discordgo.Session

func InitializeDiscord(token string, appID string, guildID string) *discordgo.Session {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating discord session: %v", err)
	}
	session.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages
	BotInstance = session
	return session
}

func GetCurrentBot() *discordgo.Session {
	return BotInstance
}

func OpenDiscordConnection(session *discordgo.Session) error {
	err := session.Open()
	if err != nil {
		return err
	}
	log.Println("Discord connection opened")
	return nil
}

func CloseDiscord(session *discordgo.Session) {
	session.Close()
	log.Println("Discord connection closed")
}
