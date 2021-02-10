package app

import (
	"net/http"

	"github.com/Kungfucoding23/microservices-go/mvc/controllers"
)

// StartApp ...
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
