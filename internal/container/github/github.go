package github

import "context"

type GithubClient interface {
	ListBranch(ctx context.Context)
	HookAction(ctx context.Context, target string)
}
