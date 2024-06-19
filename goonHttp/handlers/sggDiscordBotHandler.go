package handlers

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"

	"github.com/kjblanchard/sggDiscordBot/discord/webhookReactions"
)

type PushEventPayload struct {
	Ref        string `json:"ref"`
	Repository struct {
		Name string `json:"name"`
	} `json:"repository"`
	// Add other fields as needed based on the event payload
}
type ReleaseEventPayload struct {
	Action  string `json:"action"`
	Release struct {
		URL         string `json:"url"`
		AssetsURL   string `json:"assets_url"`
		UploadURL   string `json:"upload_url"`
		HTMLURL     string `json:"html_url"`
		ID          int    `json:"id"`
		TagName     string `json:"tag_name"`
		Name        string `json:"name"`
		CreatedAt   string `json:"created_at"`
		PublishedAt string `json:"published_at"`
		Body        string `json:"body"`
	} `json:"release"`
	Repository struct {
		Url string `json:"html_url"`
	} `json:"repository"`
}

const secret = "kevinb"

func HandleSupergoonGamesDiscordBot(w http.ResponseWriter, r *http.Request) {

	// Only handle Post requests.
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Verify the payload signature
	signature := r.Header.Get("X-Hub-Signature")
	if !verifySignature(signature, body) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	eventType := r.Header.Get("X-GitHub-Event")
	if eventType == "release" {
		// We should handle the release and deploy it
		var payload ReleaseEventPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			http.Error(w, "Failed to parse JSON payload", http.StatusBadRequest)
			return
		}
		if payload.Action == "published" {
			webhookReactions.PostNewRelease(payload.Repository.Url, payload.Release.HTMLURL, payload.Release.Name, payload.Release.Body, payload.Release.TagName)
		}

	}

	// Respond with a success status
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
