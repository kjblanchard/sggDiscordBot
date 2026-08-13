package discord_bot

import (
	"log"

	"github.com/bwmarrin/discordgo"

	"sgg-discord-bot/discord_bot/slash_commands"
)

func AddAllSlashCommands(session *discordgo.Session, guildID string) {
	commands := []*discordgo.ApplicationCommand{
		slash_commands.HelloWorldCommand,
		slash_commands.CheckReposCommand,
	}

	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}
		switch i.ApplicationCommandData().Name {
		case "hello-world":
			slash_commands.HandleHelloWorld(s, i)
		case "check-repos":
			slash_commands.HandleCheckRepos(s, i)
		}
	})

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		for _, cmd := range commands {
			_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, cmd)
			if err != nil {
				log.Printf("Error creating command %s: %v", cmd.Name, err)
			}
		}
		log.Println("Slash commands synced")
	})
}
