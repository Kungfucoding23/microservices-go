package app

import (
	"github.com/Kungfucoding23/microservices-go/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
