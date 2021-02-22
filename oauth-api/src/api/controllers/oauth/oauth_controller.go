package oauth

import (
	"net/http"

	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/utils/errors"
	"github.com/Kungfucoding23/microservices-go/oauth-api/src/api/domain/oauth"
	"github.com/Kungfucoding23/microservices-go/oauth-api/src/api/services"
	"github.com/gin-gonic/gin"
)

//CreateAccessToken creates the token
func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, token)
}

//GetAccessToken ...
func GetAccessToken(c *gin.Context) {
	token, err := services.OauthService.GetAccessToken(c.Param("token_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, token)
}
