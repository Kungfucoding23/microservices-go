package services

import (
	"github.com/Kungfucoding23/microservices-go/mvc/domain"
	"github.com/Kungfucoding23/microservices-go/mvc/utils"
)

type usersService struct{}

var (
	//UsersService struct
	UsersService usersService
)

// GetUser service
func (u *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
