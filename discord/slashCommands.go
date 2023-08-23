package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kjblanchard/sggDiscordBot/discord/slashCommands"
	"log"
)


func AddAllSlashCommands() {
	var allApplicationCommands []*discordgo.ApplicationCommand
	allApplicationCommands = slashCommands.AddHelloWorldSlashCommand(s, allApplicationCommands)
	allApplicationCommands = slashCommands.AddCheckReposSlashCommand(s, allApplicationCommands)
	_, err := s.ApplicationCommandBulkOverwrite(discordApplicationId, supergoonGamesServerId, allApplicationCommands)
	if err != nil {
		log.Fatal("Error adding application commands", err)
	}

}
