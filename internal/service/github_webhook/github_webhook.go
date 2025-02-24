package github_webhook

import (
	"context"
	"encoding/json"
	"fmt"

	"connectrpc.com/connect"
	github_webhookv1 "github.com/gandelm/gandelm/generated/protocol/github_webhook/v1"
	"github.com/gandelm/gandelm/generated/protocol/github_webhook/v1/github_webhookv1connect"
	"github.com/gandelm/gandelm/internal/container"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ github_webhookv1connect.GitHubWebhookServiceHandler = &GithubWebhookService{}

func NewGithubWebhookService(container container.Containerer) *GithubWebhookService {
	return &GithubWebhookService{container: container}
}

type GithubWebhookService struct {
	container container.Containerer
}

func (g *GithubWebhookService) HandleWebhook(ctx context.Context, req *connect.Request[github_webhookv1.GitHubWebhookRequest]) (*connect.Response[emptypb.Empty], error) {
	eventType := req.Header().Get("X-GitHub-Event") // GitHub Webhook のイベント種別
	header := req.Header()

	fmt.Println("Received GitHub event:", eventType)
	fmt.Println("Received header:", header)

	jsonData, _ := json.Marshal(req.Msg.Payload)
	fmt.Printf("Payload: %s\n", jsonData)

	return connect.NewResponse(&emptypb.Empty{}), nil
}
