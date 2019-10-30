package github

import (
	"context"
	"github.com/Thiamath/repo2graph/entities"
	"github.com/google/go-github/github"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

func GetNewClientFromToken(ghToken string) *github.Client {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: ghToken})
	return github.NewClient(oauth2.NewClient(context.Background(), tokenSource))
}

func CraftNodes(repositories []*github.Repository) (nodes []*entities.Node) {
	nodes = make([]*entities.Node, len(repositories))
	for ix, repository := range repositories {
		nodes[ix] = &entities.Node{
			Id:    repository.GetFullName(),
			Label: repository.GetName(),
		}
	}
	return nodes
}

// GetRepositories Retrieves all the repos from a token
func GetRepositories(ghClient *github.Client, ctx context.Context) (repositories []*github.Repository, err error) {
	options := github.RepositoryListOptions{
		Affiliation: "organization_member",
		ListOptions: github.ListOptions{PerPage: 25},
	}
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

func GetRepoContents(repository github.Repository) (content []*github.RepositoryContent) {

}
