package oauth

import (
	"strings"

	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/utils/errors"
)

//AccessTokenRequest struct
type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Validate validates the user
func (r *AccessTokenRequest) Validate() errors.APIError {
	r.Username = strings.TrimSpace(r.Username)
	if r.Username == "" {
		return errors.NewBadRequestError("invalid username")
	}

	r.Password = strings.TrimSpace(r.Password)
	if r.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
