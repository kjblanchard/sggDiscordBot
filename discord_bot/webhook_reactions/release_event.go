package webhook_reactions

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"

	discordbot "sgg-discord-bot/discord_bot"
)

func PostNewRelease(url string, releaseURL string, releaseName string, releaseBody string, tagName string) {
	bot := discordbot.GetCurrentBot()
	if bot == nil {
		log.Println("Bot instance not set")
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "A new release has just been posted",
		Description: fmt.Sprintf("Check out the latest release for Supergoon RPG with tag %s\nPlay the emscripten build here https://escapethefate-dev.supergoon.com", tagName),
		Color:       0x00FF00,
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Release URL (downloads and notes)", Value: releaseURL, Inline: true},
			{Name: "Name", Value: releaseName, Inline: true},
			{Name: "Release Body", Value: releaseBody, Inline: false},
			{Name: "Repository Url", Value: url, Inline: false},
			{Name: "Post issues here.", Value: fmt.Sprintf("%s/issues", url), Inline: false},
		},
	}

	_, err := bot.ChannelMessageSendEmbed(RPGNotificationsChannelID, embed)
	if err != nil {
		log.Printf("Error sending release message: %v", err)
	}
}
