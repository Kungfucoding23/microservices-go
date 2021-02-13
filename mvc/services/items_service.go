package services

import (
	"net/http"

	"github.com/Kungfucoding23/microservices-go/mvc/domain"
	"github.com/Kungfucoding23/microservices-go/mvc/utils"
)

func getItem(itemID string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
