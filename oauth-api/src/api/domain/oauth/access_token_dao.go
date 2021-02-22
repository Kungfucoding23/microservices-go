package oauth

import (
	"fmt"

	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/utils/errors"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

//Save at db
func (at *AccessToken) Save() errors.APIError {
	at.AccessToken = fmt.Sprintf("USR_%d", at.UserID)
	tokens[at.AccessToken] = at
	return nil
}

//GetAccessTokenByToken ..
func GetAccessTokenByToken(accessToken string) (*AccessToken, errors.APIError) {
	token := tokens[accessToken]
	if token == nil {
		return nil, errors.NewNotFoundError("no access token found with given parameters")
	}
	return token, nil
}
