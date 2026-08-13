package goon_http

import (
	"log"
	"net/http"

	"sgg-discord-bot/goon_http/handlers"
)

func StartServer() {
	mux := http.NewServeMux()
	handlers.RegisterHandlers(mux)

	log.Println("Starting server..")
	go func() {
		if err := http.ListenAndServe(":80", mux); err != nil {
			log.Printf("HTTP server error: %v", err)
		}
	}()
}
