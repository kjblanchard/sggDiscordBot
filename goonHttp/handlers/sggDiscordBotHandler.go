package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

type PushEventPayload struct {
	Ref        string `json:"ref"`
	Repository struct {
		Name string `json:"name"`
	} `json:"repository"`
	// Add other fields as needed based on the event payload
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

	// Parse the JSON payload into a struct
	var payload PushEventPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Failed to parse JSON payload", http.StatusBadRequest)
		return
	}

	// Process the payload data
	fmt.Printf("Received push event on repository %s, ref: %s\n", payload.Repository.Name, payload.Ref)

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
