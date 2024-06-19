package slashCommands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kjblanchard/sggDiscordBot/goonGithub"
	"log"
)

var (
	sysopRoleId = "907314874101665823"
)

func CheckIfUserInRole(roleToCheckFor string, roles []string) bool {
	for _, role := range roles {
		if role == roleToCheckFor {
			return true
		}
	}
	return false
}

func AddCheckReposSlashCommand(session *discordgo.Session, commands []*discordgo.ApplicationCommand) []*discordgo.ApplicationCommand {
	session.AddHandler(func(
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
			startNumber := 0
			if option, ok := getOptionValue(data.Options, "start-number"); ok {
				startNumber = int(option.IntValue())
			}
			err := s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
				},
			)
			if err != nil {
				log.Print("Error defering", err)
			}
			go getReposAndRespondToInteraction(i.Interaction, session, startNumber)
		}
	})

	command := &discordgo.ApplicationCommand{
		Name:        "check-repos",
		Description: "Gets a list of 10 of the repos by kjb, starting at the first",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "start-number",
				Description: "The number to start listing repositories from",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    false,
			},
		},
	}
	return append(commands, command)
}

func getOptionValue(options []*discordgo.ApplicationCommandInteractionDataOption, name string) (*discordgo.ApplicationCommandInteractionDataOption, bool) {
	for _, option := range options {
		if option.Name == name {
			return option, true
		}
	}
	return nil, false
}

func getReposAndRespondToInteraction(interaction *discordgo.Interaction, session *discordgo.Session, startNumber int) {
	repos := goonGithub.GetAllRepos()
	field := []*discordgo.MessageEmbedField{}
	start := startNumber
	end := start + 10

	for i := start; i < end; i++ {
		repo := repos[i]
		newField := &discordgo.MessageEmbedField{
			Name:  *repo.Name,
			Value: *repo.DefaultBranch,
		}
		field = append(field, newField)
	}
	// }
	content := "Heres a list of 10 the Repos!"
	embed := &discordgo.MessageEmbed{
		Title:  "Repos",
		Fields: field,
	}
	embeds := []*discordgo.MessageEmbed{embed}
	_, err := session.InteractionResponseEdit(interaction, &discordgo.WebhookEdit{
		Content: &content,
		Embeds:  &embeds,
	})
	if err != nil {
		log.Print("Error responding", err)
	}

}
