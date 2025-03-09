package github

import (
	"context"
	"encoding/json"

	"github.com/gandelm/gandelm/internal/container"
	"github.com/google/go-github/v69/github"
	"golang.org/x/oauth2"
)

var _ container.Github = &Github{}

type Github struct {
	githubOrg  string
	githubRepo string
	token      string
}

type RepositoryDipatchMessage struct {
	Artifacts map[string]string `json:"artifacts"`
}

func (c *Github) Organization() string {
	return c.githubOrg
}

func (c *Github) Repository() string {
	return c.githubRepo
}

func (c *Github) Client(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: c.token})
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func (c *Github) HookAction(ctx context.Context, eventType string, input map[string]string) (any, error) {
	message, err := json.Marshal(RepositoryDipatchMessage{
		Artifacts: input,
	})
	payload := json.RawMessage(message)
	if err != nil {
		return nil, err
	}

	_, resp, err := c.Client(ctx).Repositories.Dispatch(ctx, c.Organization(), c.Repository(), github.DispatchRequestOptions{
		EventType:     eventType,
		ClientPayload: &payload,
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Github) ListBranch(ctx context.Context) ([]string, error) {
	branches, _, err := c.Client(ctx).Repositories.ListBranches(ctx, c.Organization(), c.Repository(), nil)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(branches))
	for _, branch := range branches {
		result = append(result, branch.GetName())
	}

	return result, nil
}

func NewGithub(org, repo string) *Github {
	return &Github{githubOrg: org, githubRepo: repo}
}
