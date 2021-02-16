package repositories

import (
	"net/http"

	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/services"
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/utils/errors"

	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/domain/repo"
	"github.com/gin-gonic/gin"
)

//CreateRepocontroller controller
func CreateRepocontroller(c *gin.Context) {
	var request repo.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	// clientId := c.GetHeader("X-Client-Id")

	result, err := services.RepoService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	//return the result
	c.JSON(http.StatusCreated, result)
}