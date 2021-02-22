package oauth

import (
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/utils/errors"
)

const (
	queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User{
		"ale": {ID: 123, Username: "Ale"},
	}
)

//GetUserByUsernameAndPassword ..
func GetUserByUsernameAndPassword(username string, password string) (*User, errors.APIError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError("no user found with given parameters")
	}

	return user, nil
}
