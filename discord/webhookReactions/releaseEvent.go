package webhookReactions

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kjblanchard/sggDiscordBot/discord"
)

func PostNewRelease(url string, releaseUrl string, releaseName string, releaseBody string, tagName string) {

	embed := &discordgo.MessageEmbed{
		Title:       "A new release has just been posted",
		Description: fmt.Sprintf("Check out the latest release for Supergoon RPG with tag %s\nPlay the emscripten build here https://escapethefate.supergoon.com or the dev build here https://escapethefate-dev.supergoon.com", tagName),
		Color:       0x00ff00, // Green color
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "Release URL (downloads and notes)",
				Value:  releaseUrl,
				Inline: true,
			},
			{
				Name: "Name",
				Value:  releaseName,
				Inline: true,
			},
			{
				Name: "Release Body",
				Value:  releaseBody,
				Inline: false,
			},
			{
				Name: "Repository Url",
				Value:  url,
				Inline: false,
			},
			{
				Name: "Post issues here.",
				Value:  fmt.Sprintf("%s/issues", url),
				Inline: false,
			},
		},
	}

	// Send the embed message to the specified channel
	_, err := discord.S.ChannelMessageSendEmbed(rpgNotificationsChannelId, embed)
	if err != nil {
		fmt.Println("Error sending message,", err)
		return
	}
	// fmt.Println(message)

}
