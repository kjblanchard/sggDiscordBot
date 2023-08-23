package slashCommands
import(
	"github.com/bwmarrin/discordgo"
	"log"
)

func AddHelloWorldSlashCommand(session *discordgo.Session,  commands []*discordgo.ApplicationCommand) []*discordgo.ApplicationCommand {
	command := &discordgo.ApplicationCommand{
		Name:        "hello-world",
		Description: "Showcase of a basic slash command",
	}
	session.AddHandler(func(
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
	commands = append(commands, command)
	return commands
}