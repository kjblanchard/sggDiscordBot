package webhookReactions

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kjblanchard/sggDiscordBot/discord"
)

func PostNewRelease(url string, releaseUrl string, releaseName string, releaseBody string, tagName string) {

	// Unmarshal the JSON data into the TestEvent struct
	embed := &discordgo.MessageEmbed{
		Title:       "A new release has just appeared",
		Description: fmt.Sprintf("Check out the release for Supergoon RPG with tag %s\nPlay the emscripten build here https://rpg.supergoon.com", tagName),
		Color:       0x00ff00, // Green color
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "Release URL (downloads and notes)",
				// Value:  testEvent.Release.HTMLURL,
				Value:  releaseUrl,
				Inline: true,
			},
			{
				Name: "Name",
				// Value:  testEvent.Release.Name,
				Value:  releaseName,
				Inline: true,
			},
			{
				Name: "Release Body",
				// Value:  testEvent.Release.Body,
				Value:  releaseBody,
				Inline: false,
			},
			{
				Name: "Repository Url",
				// Value:  testEvent.Repository.Url,
				Value:  url,
				Inline: false,
			},
			{
				Name: "Post issues here!",
				// Value:  fmt.Sprintf("%s/issues", testEvent.Repository.Url),
				Value:  fmt.Sprintf("%s/issues", url),
				Inline: false,
			},
			{
				Name:   "Project board with upcoming work and such",
				Value:  "https://github.com/users/kjblanchard/projects/11",
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
