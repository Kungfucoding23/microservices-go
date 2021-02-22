package app

import (
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/controllers/polo"
	"github.com/Kungfucoding23/microservices-go/oauth-api/src/api/controllers/oauth"
)

func mapUrls() {
	router.GET("/marco", polo.MarcoController)

	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
