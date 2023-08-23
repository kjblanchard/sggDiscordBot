package goonGithub

import (
	"context"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

var (
	githubClient *github.Client
	ctx context.Context
)

func InitializeGithub(token string) {
	ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	githubClient = github.NewClient(tc)
}
