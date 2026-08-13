package slash_commands

import (
	"log"

	"github.com/bwmarrin/discordgo"

	"sgg-discord-bot/goon_github"
)

const SysopRoleID = "907314874101665823"

var CheckReposCommand = &discordgo.ApplicationCommand{
	Name:        "check-repos",
	Description: "Gets a list of 10 of the repos by kjb, starting at the first",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionInteger,
			Name:        "start_number",
			Description: "The number to start listing repositories from",
			Required:    false,
		},
	},
}

func checkIfUserInRole(roleToCheck string, roles []string) bool {
	for _, role := range roles {
		if role == roleToCheck {
			return true
		}
	}
	return false
}

func HandleCheckRepos(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkIfUserInRole(SysopRoleID, i.Member.Roles) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Only Sysop can perform this action",
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	startNumber := int64(0)
	options := i.ApplicationCommandData().Options
	for _, opt := range options {
		if opt.Name == "start_number" {
			startNumber = opt.IntValue()
		}
	}

	repos := goon_github.GetAllRepos()
	end := int(startNumber) + 10

	var fields []*discordgo.MessageEmbedField
	for idx := int(startNumber); idx < end && idx < len(repos); idx++ {
		repo := repos[idx]
		branch := "main"
		if repo.GetDefaultBranch() != "" {
			branch = repo.GetDefaultBranch()
		}
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   repo.GetName(),
			Value:  branch,
			Inline: false,
		})
	}

	embed := &discordgo.MessageEmbed{
		Title:  "Repos",
		Fields: fields,
	}

	_, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: "Heres a list of 10 the Repos!",
		Embeds:  []*discordgo.MessageEmbed{embed},
	})
	if err != nil {
		log.Printf("Error sending followup: %v", err)
	}
}
