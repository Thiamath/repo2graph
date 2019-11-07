package handlers

import (
	"context"
	"github.com/Thiamath/repo2graph/entities"
	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"time"
)

const TOKEN string = "GITHUB_TOKEN"

func GetNewClientFromToken(ghToken string) *github.Client {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: ghToken})
	return github.NewClient(oauth2.NewClient(context.Background(), tokenSource))
}

// GetRepositories Retrieves all the repos from a token
func GetRepositories(ghToken string) (repositories []*entities.Repository, err *entities.Error) {
	ghClient := GetNewClientFromToken(ghToken)

	getReposCtx, getReposCancel := context.WithTimeout(context.Background(), time.Minute)
	defer getReposCancel()
	options := github.RepositoryListOptions{
		Affiliation: "organization_member",
		ListOptions: github.ListOptions{PerPage: 25},
	}
	for {
		pagedRepos, response, err := ghClient.Repositories.List(getReposCtx, "", &options)
		if err != nil {
			log.Error(err)
			ghErr := err.(*github.ErrorResponse)
			return nil, &entities.Error{
				Error:     true,
				ErrorCode: ghErr.Response.StatusCode,
				Message:   ghErr.Error(),
			}
		}
		repositories = append(repositories, toInternalRepositoryList(pagedRepos)...)
		if response.NextPage == 0 {
			break
		}
		options.Page = response.NextPage
	}
	return repositories, nil
}

func toInternalRepositoryList(githubRepositoryList []*github.Repository) (internalRepositoryList []*entities.Repository) {
	internalRepositoryList = make([]*entities.Repository, len(githubRepositoryList))
	for i, githubRepository := range githubRepositoryList {
		internalRepositoryList[i] = toInternalRepository(githubRepository)
	}
	return internalRepositoryList
}

func toInternalRepository(githubRepository *github.Repository) (internalRepository *entities.Repository) {
	return &entities.Repository{
		Id:   githubRepository.GetName(),
		Name: githubRepository.GetName(),
	}
}
