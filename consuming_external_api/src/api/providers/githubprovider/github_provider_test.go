package githubprovider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/client/restclient"
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/domain/github"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuth)
	assert.EqualValues(t, "token %s", headerAuthFormat)
	assert.EqualValues(t, "https://api.github.com/user/repos", urlCreateRepo)
}

func TestGetAuthHeader(t *testing.T) {
	header := getAuthHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoErrorClient(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Err:        errors.New("invalid restclient response"),
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid restclient response", err.Message)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	restclient.FlushMockups()
	invalidCloser, _ := os.Open("-asd123")
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid response body", err.Message)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
}
func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid json response body", err.Message)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
}
func TestCreateRepoUnauthorized(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{
				"message": "Requires authentication",
				"documentation_url": "https://docs.github.com/rest/reference/repos#create-a-repository-for-the-authenticated-user"
			}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Requires authentication", err.Message)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
}
func TestCreateRepoInvalidSuccessResponse(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": "123"}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "error when trying to Unmarshal github create repo response", err.Message)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
}
func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": 123,
			"name": "Golang-microservices-Tutorial-", "full_name": "Kungfucoding23/Golang-microservices-Tutorial-"}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 123, response.ID)
	assert.EqualValues(t, "Golang-microservices-Tutorial-", response.Name)
	assert.EqualValues(t, "Kungfucoding23/Golang-microservices-Tutorial-", response.FullName)
}
