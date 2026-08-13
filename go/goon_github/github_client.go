package goon_github

import (
	"context"

	"github.com/google/go-github/v60/github"
)

var GithubClient *github.Client

func InitializeGithub(token string) {
	GithubClient = github.NewClient(nil).WithAuthToken(token)
}

func GetContext() context.Context {
	return context.Background()
}
