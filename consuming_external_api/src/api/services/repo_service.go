package services

import (
	"strings"

	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/config"
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/domain/github"
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/domain/repo"
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/providers/githubprovider"
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/utils/errors"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateRepo(request repo.CreateRepoRequest) (*repo.CreateRepoResponse, errors.APIError)
}

var (
	//RepoService ..
	RepoService repoServiceInterface
)

func init() {
	RepoService = &repoService{}
}

func (service *repoService) CreateRepo(input repo.CreateRepoRequest) (*repo.CreateRepoResponse, errors.APIError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("Invalid repo name")
	}
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := githubprovider.CreateRepo(config.GetGithubToken(), request)
	if err != nil {
		return nil, errors.NewAPIError(err.StatusCode, err.Message)
	}
	result := repo.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}
