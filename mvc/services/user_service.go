package services

import (
	"github.com/Kungfucoding23/microservices-go/mvc/domain"
	"github.com/Kungfucoding23/microservices-go/mvc/utils"
)

// GetUser service
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
