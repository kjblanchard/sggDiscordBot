package github

import (
	"log"
)


func GetAllRepos() {

	repos, _, err := githubClient.Repositories.List(ctx, "kjblanchard", nil)
	if err != nil {
		log.Printf("Error listing repositories: Error %s", err)
	}

	for _, repo := range repos {
		log.Printf("Repo id is %d and name is %s ", repo.ID, *repo.Name)
	}

}
