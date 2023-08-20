package main

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"fmt"
	"os/signal"
	"strings"
)

type appSettings struct {
	Token                  string `json:"token"`
	AppId                  string `json:"appId"`
	SupergoonGamesServerId string `json:"supergoonGamesServerId"`
}

func main() {
	jsonContent, err := os.ReadFile("appsettings.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
		return
	}
	var config appSettings
	err = json.Unmarshal(jsonContent, &config)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
		return
	}
	s, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal("Error creating session: ", err)
	}
	s.AddHandler(newMessage)

	 // open session
	 err = s.Open()
	 if err != nil {
		log.Fatal("Error opening websocket connection: ", err)
	 }
	 defer s.Close() // close session, after function termination
	 fmt.Println("Bot running....")
	 c := make(chan os.Signal, 1)
	 signal.Notify(c, os.Interrupt)
	 <-c
}

	func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

		/* prevent bot responding to its own message
		this is achived by looking into the message author id
		if message.author.id is same as bot.author.id then just return
		*/
		if message.Author.ID == discord.State.User.ID {
		 return
		}

		// respond to user message if it contains `!help` or `!bye`
		switch {
		case strings.Contains(message.Content, "!help"):
		 discord.ChannelMessageSend(message.ChannelID, "Hello World😃")
		case strings.Contains(message.Content, "!bye"):
		 discord.ChannelMessageSend(message.ChannelID, "Good Bye👋")
		 // add more cases if required
		}

}
