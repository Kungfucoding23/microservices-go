package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Ale", LastName: "Santin", Email: "alejandro@gmail.com"},
	}
)

// GetUser from domain
func GetUser(userID int64) (*User, error) {
	//here we connect to the db (using users map as db)
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, errors.New(fmt.Sprintf("User %v was not found", userID))
}
