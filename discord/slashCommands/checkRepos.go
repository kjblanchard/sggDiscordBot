package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func AddCheckReposSlashCommand(session *discordgo.Session, commands []*discordgo.ApplicationCommand) []*discordgo.ApplicationCommand {
	command := &discordgo.ApplicationCommand{
		Name:        "check-repos",
		Description: "Checks to see all of the repos available",
	}
	command = append(command, command)
	s.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		data := i.ApplicationCommandData()
		if data.Name != "check-repos" {
			return
		}
		roles := i.Member.Roles
		if !CheckIfUserInRole(sysopRoleId, roles) {
			err := s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Data: &discordgo.InteractionResponseData{
						Content: "Only Sysop can perform this action",
					},
				},
			)
			if err != nil {
				log.Print("Error responding! ", err)
			}
		} else {
			err := s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
				},
			)
			if err != nil {
				log.Print("Error defering", err)
			}
			go getReposAndRespondToInteraction(i.Interaction)
		}
	})
}

func getReposAndRespondToInteraction(interaction *discordgo.Interaction) {
	repos := goonGithub.GetAllRepos()
	field := []*discordgo.MessageEmbedField{}
	for _, repo := range repos {
		newField := &discordgo.MessageEmbedField{
			Name:  *repo.Name,
			Value: *repo.DefaultBranch,
		}
		field = append(field, newField)
	}

	content := "Heres a list of all the Repos!"
	embed := &discordgo.MessageEmbed{
		Title:  "Repos",
		Fields: field,
	}
	embeds := []*discordgo.MessageEmbed{embed}
	_, err := s.InteractionResponseEdit(interaction, &discordgo.WebhookEdit{
		Content: &content,
		Embeds:  &embeds,
	})
	if err != nil {
		log.Print("Error responding", err)
	}

}
