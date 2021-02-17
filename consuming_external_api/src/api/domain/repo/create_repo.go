package repo

import (
	"strings"

	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/utils/errors"
)

//CreateRepoRequest ..
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

//Validate the repo
func (r *CreateRepoRequest) Validate() errors.APIError {
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return errors.NewBadRequestError("invalid repository name")
	}
	return nil
}

//CreateRepoResponse ..
type CreateRepoResponse struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

//CreateReposResponse struct
type CreateReposResponse struct {
	StatusCode int                 `json:"status"`
	Results    []CreateReposResult `json:"results"`
}

//CreateReposResult struct
type CreateReposResult struct {
	Response *CreateRepoResponse `json:"repo"`
	Error    errors.APIError     `json:"error"`
}
