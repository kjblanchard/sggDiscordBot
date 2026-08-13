package goon_github

import (
	"log"

	"github.com/google/go-github/v60/github"
)

func GetAllRepos() []*github.Repository {
	if GithubClient == nil {
		log.Println("GitHub client not initialized")
		return nil
	}
	repos, _, err := GithubClient.Repositories.List(GetContext(), "kjblanchard", nil)
	if err != nil {
		log.Printf("Error listing repositories: %v", err)
		return nil
	}
	return repos
}
