package services

import (
	"net/http"
	"sync"

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
	CreateRepos(request []repo.CreateRepoRequest) (repo.CreateReposResponse, errors.APIError)
}

var (
	//RepoService ..
	RepoService repoServiceInterface
)

func init() {
	RepoService = &repoService{}
}

func (service *repoService) CreateRepo(input repo.CreateRepoRequest) (*repo.CreateRepoResponse, errors.APIError) {

	if err := input.Validate(); err != nil {
		return nil, err
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

func (service *repoService) CreateRepos(requests []repo.CreateRepoRequest) (repo.CreateReposResponse, errors.APIError) {
	//Here we have an array, so we need every input request to be a valid request or we want to valid and process them in a concurrent way and delegating the responsability of taking this different request and handle this in a separated way.
	input := make(chan repo.CreateReposResult)
	output := make(chan repo.CreateReposResponse)

	defer close(output)

	var wg sync.WaitGroup
	//wg is a control mechanism that we have in order to block the ejecution until the work is done

	go service.handleRepoResults(&wg, input, output)
	//n requests to process
	for _, current := range requests {
		//for each request that we have in the slice we create a go routine that will handle them separated in a concurrent way
		wg.Add(1)
		go service.CreateRepoConcurrent(current, input)
	}
	wg.Wait() //the execution will freeze in here until the wg reaches zero
	close(input)

	result := <-output

	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations++
		}
	}
	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}
	return result, nil
}

func (service *repoService) handleRepoResults(wg *sync.WaitGroup, input chan repo.CreateReposResult, output chan repo.CreateReposResponse) {
	var results repo.CreateReposResponse
	for incomingEvent := range input {
		repoResult := repo.CreateReposResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}
	output <- results
}

func (service *repoService) CreateRepoConcurrent(input repo.CreateRepoRequest, output chan repo.CreateReposResult) {
	if err := input.Validate(); err != nil {
		output <- repo.CreateReposResult{Error: err}
		return
	}
	result, err := service.CreateRepo(input)
	if err != nil {
		output <- repo.CreateReposResult{Error: err}
		return
	}
	output <- repo.CreateReposResult{Response: result}
}
