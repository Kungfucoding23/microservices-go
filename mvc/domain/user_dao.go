package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kungfucoding23/microservices-go/mvc/utils"
)

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Ale", LastName: "Santin", Email: "alejandro@gmail.com"},
	}
	//UserDao struct
	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

// GetUser from domain
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	log.Println("We´re accessing the database")
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v does not exists", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
