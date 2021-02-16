package app

import (
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/controllers/polo"
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.MarcoController)
	router.POST("/repositories", repositories.CreateRepocontroller)
}
