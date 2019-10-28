package github

import (
	"context"
	"github.com/Thiamath/repo2graph/entities"
	"github.com/google/go-github/github"
	"github.com/rs/zerolog/log"
)

func CraftNodes(repositories []*github.Repository) []*entities.Node {
	nodes := make([]*entities.Node, 0)
	for _, repository := range repositories {
		nodes = append(nodes, &entities.Node{
			Id:    repository.GetFullName(),
			Label: repository.GetName(),
		})
	}
	return nodes
}

// GetRepositories Retrieves all the repos from a token
func GetRepositories(ghClient *github.Client, ctx context.Context) ([]*github.Repository, error) {
	options := github.RepositoryListOptions{
		Affiliation: "organization_member",
		ListOptions: github.ListOptions{PerPage: 25},
	}
	var repositories []*github.Repository
	for {
		pagedRepos, response, err := ghClient.Repositories.List(ctx, "", &options)
		if err != nil {
			log.Error().Err(err).Msg("error listing repos")
			return nil, err
		}
		repositories = append(repositories, pagedRepos...)
		if response.NextPage == 0 {
			break
		}
		options.Page = response.NextPage
	}
	return repositories, nil
}
