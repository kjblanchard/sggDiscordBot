package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kjblanchard/sggDiscordBot/github"
	"log"
)

var (
	s *discordgo.Session
	discordApplicationId string
	supergoonGamesServerId string
)

func InitializeDiscord(token string, appId string, supergoonServerId string) {
	s, _ = discordgo.New("Bot " + token)
	discordApplicationId = appId
	supergoonGamesServerId = supergoonServerId
}

func AddAllSlashCommands() {
	_, err := s.ApplicationCommandBulkOverwrite(discordApplicationId, supergoonGamesServerId, []*discordgo.ApplicationCommand{
		{
			Name:        "check-repos",
			Description: "Checks to see all of the repos available",
		},
		{
			Name:        "hello-world",
			Description: "Showcase of a basic slash command",
		},
	})
	if err != nil {
		log.Fatal("Error doing the thing", err)
		// Handle the error
	}
	s.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		data := i.ApplicationCommandData()
		switch data.Name {
		case "hello-world":
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
				log.Fatal("Error registering", err)
				// Handle the error
			}
		case "check-repos":
			roles := i.Member.Roles

			github.GetAllRepos()

			for _, role := range roles {
				log.Print(role)
			}
			err := s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Heres a list of all of Kevins Repos..",
						Embeds: []*discordgo.MessageEmbed{
							{
								Title: "Repo",
							},
						},
					},
				},
			)
			if err != nil {
				log.Fatal("Error registering", err)
				// Handle the error
			}
		}
	})

}

func OpenDiscordWebsocketConnection() {
	err := s.Open()
	if err != nil {
		log.Fatal("Error Opening websocket connection for discord bot!\nError: ", err)
	}

}

func CloseDiscord() {
	err := s.Close()
	if err != nil {
		log.Fatal("Error Closing bot", err)
	}
}
