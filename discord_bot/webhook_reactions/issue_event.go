package webhook_reactions

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"

	discordbot "sgg-discord-bot/discord_bot"
)

const IssueChannelID = "1536884393870893137"

func handleIssueCreateEvent(issueData map[string]any) *discordgo.MessageEmbed {
	url, _ := issueData["repository_url"].(string)
	issueNum, _ := issueData["number"].(float64)
	title, _ := issueData["title"].(string)
	userData, _ := issueData["user"].(map[string]any)
	createdBy, _ := userData["login"].(string)
	body, _ := issueData["body"].(string)

	bodyText := body
	if len(body) >= 1200 {
		bodyText = "Body too long for discord, view url"
	}

	embed := &discordgo.MessageEmbed{
		Title:       "New issue created",
		Description: fmt.Sprintf("#%d: %s:\n%s", int(issueNum), title, bodyText),
		Color:       0x2784F5,
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Issue Url", Value: url, Inline: false},
			{Name: "Created by", Value: createdBy, Inline: false},
		},
	}
	return embed
}

func handleIssueClosedEvent(issueData map[string]any) *discordgo.MessageEmbed {
	url, _ := issueData["url"].(string)
	issueNum, _ := issueData["number"].(float64)
	title, _ := issueData["title"].(string)
	userData, _ := issueData["user"].(map[string]any)
	closedBy, _ := userData["login"].(string)

	embed := &discordgo.MessageEmbed{
		Title:       "Issue closed",
		Description: fmt.Sprintf("#%d: %s", int(issueNum), title),
		Color:       0xE727F5,
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Issue Url", Value: url, Inline: false},
			{Name: "Closed by", Value: closedBy, Inline: false},
		},
	}
	return embed
}

func HandleIssueEvent(payload map[string]any) {
	bot := discordbot.GetCurrentBot()
	if bot == nil {
		log.Println("Current bot not set, not posting issue event")
		return
	}
	issueType, ok := payload["action"].(string)
	if !ok {
		log.Println("Issue type could not be determined!")
		return
	}

	issueData, ok := payload["issue"].(map[string]any)
	if !ok {
		log.Println("Could not parse issue data")
		return
	}
	var embed *discordgo.MessageEmbed
	switch issueType {
	case "closed":
		embed = handleIssueClosedEvent(issueData)
	case "opened":
		embed = handleIssueCreateEvent(issueData)
	}

	_, err := bot.ChannelMessageSendEmbed(IssueChannelID, embed)
	if err != nil {
		log.Printf("Error sending issue message: %v", err)
	}
}
