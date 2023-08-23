package goonGithub

import (
	"log"
	"github.com/google/go-github/v39/github"

)


func GetAllRepos() []*github.Repository {

	repos, _, err := githubClient.Repositories.List(ctx, "kjblanchard", nil)
	if err != nil {
		log.Printf("Error listing repositories: Error %s", err)
	}
	return repos
}
