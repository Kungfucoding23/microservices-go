package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/client/restclient"
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/domain/repo"
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidName(t *testing.T) {
	request := repo.CreateRepoRequest{}

	result, err := RepoService.CreateRepo("client_id", request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "invalid repository name", err.Message())
}
func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{
				"message": "Requires authentication",
				"documentation_url": "https://developer.github.com/docs"}`)),
		},
	})
	request := repo.CreateRepoRequest{Name: "testing"}

	result, err := RepoService.CreateRepo("client_id", request)

	assert.Nil(t, result)
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires authentication", err.Message())

}
func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login": "Kungfucoding23"}}`)),
		},
	})
	request := repo.CreateRepoRequest{Name: "testing"}

	result, err := RepoService.CreateRepo("client_id", request)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.EqualValues(t, 123, result.ID)
	assert.EqualValues(t, "testing", result.Name)
	assert.EqualValues(t, "Kungfucoding23", result.Owner)
}

func TestCreateRepoConcurrentInvalidRequest(t *testing.T) {
	request := repo.CreateRepoRequest{}

	output := make(chan repo.CreateReposResult)

	service := repoService{}
	go service.CreateRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Error.Message())
}
func TestCreateRepoConcurrentErrorFromGithub(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{
				"message": "Requires authentication",
				"documentation_url": "https://developer.github.com/docs"}`)),
		},
	})
	request := repo.CreateRepoRequest{Name: "testing"}

	output := make(chan repo.CreateReposResult)

	service := repoService{}
	go service.CreateRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusUnauthorized, result.Error.Status())
	assert.EqualValues(t, "Requires authentication", result.Error.Message())
}
func TestCreateRepoConcurrentNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login": "Kungfucoding23"}}`)),
		},
	})
	request := repo.CreateRepoRequest{Name: "testing"}

	output := make(chan repo.CreateReposResult)

	service := repoService{}
	go service.CreateRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Response)
	assert.EqualValues(t, 123, result.Response.ID)
	assert.EqualValues(t, "testing", result.Response.Name)
	assert.EqualValues(t, "Kungfucoding23", result.Response.Owner)
}

func TestHandleRepoResults(t *testing.T) {
	input := make(chan repo.CreateReposResult)
	output := make(chan repo.CreateReposResponse)

	var wg sync.WaitGroup

	service := repoService{}

	go service.handleRepoResults(&wg, input, output)

	wg.Add(1)
	go func() {
		input <- repo.CreateReposResult{
			Error: errors.NewBadRequestError("ivalid repository name"),
		}
	}()
	wg.Wait()
	close(input)
	results := <-output

	assert.NotNil(t, results)

	assert.EqualValues(t, 0, results.StatusCode)
	assert.EqualValues(t, 1, len(results.Results))
	assert.NotNil(t, results.Results[0].Error)
	assert.EqualValues(t, http.StatusBadRequest, results.Results[0].Error.Status())
	assert.EqualValues(t, "ivalid repository name", results.Results[0].Error.Message())

}

func TestCreateReposInvalidRequest(t *testing.T) {
	requests := []repo.CreateRepoRequest{
		{},
		{Name: "   "},
	}

	results, err := RepoService.CreateRepos(requests)

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Nil(t, results.Results[0].Response)
	assert.EqualValues(t, http.StatusBadRequest, results.StatusCode)
	assert.EqualValues(t, 2, len(results.Results))
	assert.EqualValues(t, http.StatusBadRequest, results.Results[0].Error.Status())
	assert.EqualValues(t, "invalid repository name", results.Results[0].Error.Message())

	assert.Nil(t, results.Results[1].Response)
	assert.EqualValues(t, http.StatusBadRequest, results.Results[1].Error.Status())
	assert.EqualValues(t, "invalid repository name", results.Results[1].Error.Message())
}

func TestCreateReposPartialRequest(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login": "Kungfucoding23"}}`)),
		},
	})
	requests := []repo.CreateRepoRequest{
		{},
		{Name: "testing"},
	}

	results, err := RepoService.CreateRepos(requests)

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.EqualValues(t, 2, len(results.Results))
	assert.EqualValues(t, http.StatusPartialContent, results.StatusCode)

	for _, result := range results.Results {
		if result.Error != nil {
			assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
			assert.EqualValues(t, "invalid repository name", result.Error.Message())
			continue
		}
		assert.EqualValues(t, 123, result.Response.ID)
		assert.EqualValues(t, "testing", result.Response.Name)
		assert.EqualValues(t, "Kungfucoding23", result.Response.Owner)
	}

}

func TestCreateReposWithSameName(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login": "Kungfucoding23"}}`)),
		},
	})
	requests := []repo.CreateRepoRequest{
		{Name: "testing"},
		{Name: "testing"},
	}

	results, err := RepoService.CreateRepos(requests)

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.EqualValues(t, 2, len(results.Results))

	for _, result := range results.Results {
		if result.Error != nil {
			assert.EqualValues(t, http.StatusInternalServerError, result.Error.Status())
			assert.EqualValues(t, "error when trying to Unmarshal github create repo response", result.Error.Message())
			continue
		}
		assert.EqualValues(t, 123, result.Response.ID)
		assert.EqualValues(t, "testing", result.Response.Name)
		assert.EqualValues(t, "Kungfucoding23", result.Response.Owner)
	}
}

func TestCreateReposRepoAlreadyExistsFailure(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login": "Kungfucoding23"}}`)),
		},
	})

	requests := []repo.CreateRepoRequest{
		{Name: "testing"},
		{Name: "testing"},
	}

	result, err := RepoService.CreateRepos(requests)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusPartialContent, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	for _, result := range result.Results {
		if result.Error != nil {
			assert.EqualValues(t, http.StatusInternalServerError, result.Error.Status())
			assert.EqualValues(t, "error when trying to Unmarshal github create repo response", result.Error.Message())
			continue
		}

		assert.EqualValues(t, 123, result.Response.ID)
		assert.EqualValues(t, "testing", result.Response.Name)
		assert.EqualValues(t, "Kungfucoding23", result.Response.Owner)
	}
}
