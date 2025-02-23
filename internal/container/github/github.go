package github

import (
	"context"
	"fmt"
	"strings"

	"github.com/gandelm/gandelm/internal/container"
	"github.com/google/go-github/v69/github"
)

var _ container.Github = &Github{}

type Github struct {
	githubOrg  string
	githubRepo string
}

func (c *Github) Organization() string {
	return c.githubOrg
}

func (c *Github) Repository() string {
	return c.githubRepo
}

func (g *Github) HookAction(ctx context.Context, target string) {
	panic("unimplemented")
}

func (g *Github) ListBranch(ctx context.Context) ([]string, error) {
	client := github.NewClient(nil)
	branches, _, err := client.Repositories.ListBranches(ctx, g.Organization(), g.Repository(), nil)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(branches))
	for _, branch := range branches {
		result = append(result, branch.GetName())
	}

	return result, nil
}

func NewGithub(url string) *Github {
	org, repo, err := extractGitHubRepoInfo("https://github.com/gandelm/gandelm")
	if err != nil {
		return nil
	}
	return &Github{githubOrg: org, githubRepo: repo}
}

func extractGitHubRepoInfo(url string) (string, string, error) {
	parts := strings.Split(strings.TrimPrefix(url, "https://github.com/"), "/")
	if len(parts) != 2 || len(parts) != 3 {
		return "", "", fmt.Errorf("invalid GitHub URL")
	}
	return parts[0], parts[1], nil
}
