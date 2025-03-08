package github_webhook_handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var githubSecret = os.Getenv("GITHUB_WEBHOOK_SECRET")

type GitHubEvent struct {
	Action string `json:"action,omitempty"` // workflow_run 用
}

type WorkflowRunEvent struct {
	Action      string `json:"action"`
	WorkflowRun struct {
		ID         int64  `json:"id"`
		Name       string `json:"name"`
		Status     string `json:"status"`     // "completed", "in_progress", etc.
		Conclusion string `json:"conclusion"` // "success", "failure", "cancelled"
	} `json:"workflow_run"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
}

type PushEvent struct {
	Ref    string `json:"ref"`
	Before string `json:"before"`
	After  string `json:"after"`
	Pusher struct {
		Name string `json:"name"`
	} `json:"pusher"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
}

func GithubWebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	signature := r.Header.Get("X-Hub-Signature-256")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if !validateSignature(body, signature) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	eventType := r.Header.Get("X-GitHub-Event")

	switch eventType {
	case "workflow_run":
		var event WorkflowRunEvent
		if err := json.Unmarshal(body, &event); err != nil {
			http.Error(w, "Failed to parse workflow_run JSON", http.StatusBadRequest)
			return
		}
		handleWorkflowRunEvent(event)

	case "push":
		var event PushEvent
		if err := json.Unmarshal(body, &event); err != nil {
			http.Error(w, "Failed to parse push JSON", http.StatusBadRequest)
			return
		}
		handlePushEvent(event)

	default:
		fmt.Printf("❌ Ignoring event type: %s\n", eventType)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Webhook received successfully")
}

func validateSignature(payload []byte, signature string) bool {
	if githubSecret == "" {
		fmt.Println("Warning: No GitHub secret set. Skipping signature validation.")
		return true
	}
	mac := hmac.New(sha256.New, []byte(githubSecret))
	mac.Write(payload)
	expectedMAC := mac.Sum(nil)
	expectedSignature := "sha256=" + hex.EncodeToString(expectedMAC)
	return hmac.Equal([]byte(expectedSignature), []byte(signature))
}

func handleWorkflowRunEvent(event WorkflowRunEvent) {
	fmt.Printf("🔄 Workflow '%s' in repo '%s' is now %s (conclusion: %s)\n",
		event.WorkflowRun.Name, event.Repository.FullName, event.WorkflowRun.Status, event.WorkflowRun.Conclusion)

	if event.Action == "completed" {
		if event.WorkflowRun.Conclusion == "success" {
			fmt.Println("✅ Workflow succeeded!")
		} else {
			fmt.Println("❌ Workflow failed!")
		}
	}
}

func handlePushEvent(event PushEvent) {
	fmt.Printf("🚀 Push event in repo '%s' by '%s': %s → %s (branch: %s)\n",
		event.Repository.FullName, event.Pusher.Name, event.Before, event.After, event.Ref)
}
