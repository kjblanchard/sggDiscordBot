package handlers

import (
	"fmt"
	"net/http"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Welcome User! !")))
}