package slash_commands

import (
	"github.com/bwmarrin/discordgo"
)

var HelloWorldCommand = &discordgo.ApplicationCommand{
	Name:        "hello-world",
	Description: "Showcase of a basic slash command",
}

func HandleHelloWorld(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hello world!",
		},
	})
}
