package githubprovider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/client/restclient"
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/domain/github"
)

const (
	headerAuth       = "Authorization"
	headerAuthFormat = "token %s"
	urlCreateRepo    = "https://api.github.com/user/repos"
)

func getAuthHeader(accesToken string) string {
	return fmt.Sprintf(headerAuthFormat, accesToken)
}

//CreateRepo creates the repo
func CreateRepo(accesToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.ErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuth, getAuthHeader(accesToken))

	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Printf("error when trying to create new repo in github: %s ", err.Error())
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid response body",
		}
	}
	defer response.Body.Close() //this is executed once the function is returning

	if response.StatusCode > 299 {
		var errResponse github.ErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	//if we reach this point we know that we don´t have any error and the response is < 299
	//so 201 - created will be the case
	//this should have the final created repo
	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Printf("error when trying to unmarshal create repo successful response: %s ", err.Error())
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error when trying to Unmarshal github create repo response",
		}
	}

	return &result, nil
}
