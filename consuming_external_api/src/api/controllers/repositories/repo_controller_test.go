package repositories

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/client/restclient"
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/domain/repo"
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidJSONReq(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))
	c.Request = request
	CreateRepocontroller(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiError, err := errors.NewAPIErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiError)
	assert.EqualValues(t, http.StatusBadRequest, apiError.Status())
	assert.EqualValues(t, "Invalid json body", apiError.Message())

}
func TestCreateRepoErrorFromGithub(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	c.Request = request

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

	CreateRepocontroller(c)
	assert.EqualValues(t, http.StatusUnauthorized, response.Code)

	apiError, err := errors.NewAPIErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiError)
	assert.EqualValues(t, http.StatusUnauthorized, apiError.Status())
	assert.EqualValues(t, "Requires authentication", apiError.Message())

}
func TestCreateRepoNoError(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	c.Request = request

	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123}`)),
		},
	})

	CreateRepocontroller(c)
	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repo.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.ID)
	assert.EqualValues(t, "", result.Name)
	assert.EqualValues(t, "", result.Owner)
}
