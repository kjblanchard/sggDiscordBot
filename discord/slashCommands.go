package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kjblanchard/sggDiscordBot/goonGithub"
	"log"
)

var (
	allApplicationCommands []*discordgo.ApplicationCommand
)

func CheckIfUserInRole(roleToCheckFor string, roles []string) bool {
	for _, role := range roles {
		if role == roleToCheckFor {
			return true
		}
	}
	return false
}



func AddAllSlashCommands() {
	AddHelloWorldSlashCommand()
	AddCheckReposSlashCommand()
	_, err := s.ApplicationCommandBulkOverwrite(discordApplicationId, supergoonGamesServerId, allApplicationCommands)
	if err != nil {
		log.Fatal("Error adding application commands", err)
	}

}
