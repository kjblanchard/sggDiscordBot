package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kjblanchard/sggDiscordBot/github"
	"log"
)

var (
	allApplicationCommands []*discordgo.ApplicationCommand
)

func checkIfUserInRole(roleToCheckFor string, roles []string) bool {
	for _, role := range roles {
		if role == roleToCheckFor {
			return true
		}
	}
	return false
}

func AddCheckReposSlashCommand() {
	command := &discordgo.ApplicationCommand{
		Name:        "check-repos",
		Description: "Checks to see all of the repos available",
	}
	allApplicationCommands = append(allApplicationCommands, command)
	s.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		data := i.ApplicationCommandData()
		if data.Name != "check-repos" {
			return
		}
		github.GetAllRepos()
		roles := i.Member.Roles
		response := &discordgo.InteractionResponseData{}
		if !checkIfUserInRole(sysopRoleId, roles) {
			response.Content = "Only Sysop can perform this action"
		} else {
			response.Content = "Heres a list of all of Kevins Repos.."
			response.Embeds = []*discordgo.MessageEmbed{
				{
					Title: "Repo",
				},
			}
		}
		// err := s.InteractionRespond(
		// 	i.Interaction,
		// 	&discordgo.InteractionResponse{
		// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
		// 		Data: response},
		// )
		// if err != nil {
		// 	log.Fatal("Error responding", err)
		// }
		err := s.InteractionRespond(
			i.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
				Data: response},
		)
		if err != nil {
			log.Print("Error defering", err)
		}

		_ , err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &response.Content,
			Embeds: &response.Embeds,
		})
		if err != nil {
			log.Print("Error responding", err)
		}
	})
}

func AddHelloWorldSlashCommand() {
	command := &discordgo.ApplicationCommand{
		Name:        "hello-world",
		Description: "Showcase of a basic slash command",
	}
	allApplicationCommands = append(allApplicationCommands, command)
	s.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		data := i.ApplicationCommandData()
		if data.Name == "hello-world" {
			err := s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Hello world!",
					},
				},
			)
			if err != nil {
				log.Fatal("Error responding", err)
			}
		}
	})
}

func AddAllSlashCommands() {
	AddHelloWorldSlashCommand()
	AddCheckReposSlashCommand()
	_, err := s.ApplicationCommandBulkOverwrite(discordApplicationId, supergoonGamesServerId, allApplicationCommands)
	if err != nil {
		log.Fatal("Error adding application commands", err)
	}

}
