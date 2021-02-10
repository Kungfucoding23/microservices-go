package services

import "github.com/Kungfucoding23/microservices-go/mvc/domain"

// GetUser service
func GetUser(userID int64) (*domain.User, error) {
	return domain.GetUser(userID)
}
