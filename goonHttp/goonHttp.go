package goonhttp

import (
	_ "fmt"
	"log"
	"net/http"
	"github.com/kjblanchard/sggDiscordBot/goonHttp/handlers"
)

func StartServer() {
	log.Print("Starting server..")
	http.HandleFunc("/api/v1/webhooks", handlers.HandleWebhook)
	log.Fatal(http.ListenAndServe(":80", nil))
}
