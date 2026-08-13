package handlers

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"sgg-discord-bot/discord_bot/webhook_reactions"
)

const secret = "kevinb"

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/webhooks/sggDiscordBot", handleSupergoonGamesDiscordBot)
}

func handleSupergoonGamesDiscordBot(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	signature := r.Header.Get("X-Hub-Signature")
	if !verifySignature(signature, body) {
		log.Println("Invalid signature")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	eventType := r.Header.Get("X-GitHub-Event")

	switch eventType {
	case "release":
		var payload map[string]interface{}
		if err := json.Unmarshal(body, &payload); err != nil {
			break
		}
		if action, _ := payload["action"].(string); action == "published" {
			release, _ := payload["release"].(map[string]interface{})
			repository, _ := payload["repository"].(map[string]interface{})
			webhook_reactions.PostNewRelease(
				getString(repository, "html_url"),
				getString(release, "html_url"),
				getString(release, "name"),
				getString(release, "body"),
				getString(release, "tag_name"),
			)
		}
	case "issues":
		var payload map[string]interface{}
		if err := json.Unmarshal(body, &payload); err != nil {
			break
		}
		webhook_reactions.HandleIssueEvent(payload)
	}

	w.WriteHeader(http.StatusOK)
}

func verifySignature(signature string, payload []byte) bool {
	expectedMAC := calculateMAC(payload)
	return hmac.Equal([]byte(signature), []byte(expectedMAC))
}

func calculateMAC(payload []byte) string {
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write(payload)
	return "sha1=" + hex.EncodeToString(mac.Sum(nil))
}

func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}
