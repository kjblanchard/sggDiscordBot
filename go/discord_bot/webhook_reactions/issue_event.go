package webhook_reactions

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"

	discordbot "sgg-discord-bot/discord_bot"
)

const IssueChannelID = "1536884393870893137"

func HandleIssueEvent(payload map[string]interface{}) {
	bot := discordbot.GetCurrentBot()
	if bot == nil {
		log.Println("Current bot not set, not posting issue event")
		return
	}

	issueData, ok := payload["issue"].(map[string]interface{})
	if !ok {
		log.Println("Could not parse issue data")
		return
	}

	url, _ := issueData["repository_url"].(string)
	issueNum, _ := issueData["number"].(float64)
	title, _ := issueData["title"].(string)
	userData, _ := issueData["user"].(map[string]interface{})
	createdBy, _ := userData["login"].(string)
	body, _ := issueData["body"].(string)

	bodyText := body
	if len(body) >= 1200 {
		bodyText = "Body too long for discord, view url"
	}

	embed := &discordgo.MessageEmbed{
		Title:       "New issue created",
		Description: fmt.Sprintf("#%d: %s:\n%s", int(issueNum), title, bodyText),
		Color:       0x00FF00,
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Issue Url", Value: url, Inline: false},
			{Name: "Created by", Value: createdBy, Inline: false},
		},
	}

	_, err := bot.ChannelMessageSendEmbed(IssueChannelID, embed)
	if err != nil {
		log.Printf("Error sending issue message: %v", err)
	}
}
