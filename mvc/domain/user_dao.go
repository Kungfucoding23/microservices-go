package domain

import (
	"fmt"
	"net/http"

	"github.com/Kungfucoding23/microservices-go/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Ale", LastName: "Santin", Email: "alejandro@gmail.com"},
	}
)

// GetUser from domain
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	//here we connect to the db (using users map as db)
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("User %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code: "not_found",
	}
}
