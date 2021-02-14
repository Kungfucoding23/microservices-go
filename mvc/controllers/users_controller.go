package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Kungfucoding23/microservices-go/mvc/services"
	"github.com/Kungfucoding23/microservices-go/mvc/utils"
	"github.com/gin-gonic/gin"
)

// GetUser controller
func GetUser(c *gin.Context) {
	userIDParam := c.Param("user_id")
	if len(userIDParam) < 1 {
		apiErr := &utils.ApplicationError{
			Message:    "user_id is needed",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	//insert curl localhost:8080/users?user_id=123 to test
	log.Printf("about to process user_id %v", userIDParam)

	userID, err := strconv.ParseInt(userIDParam, 10, 64) //base=10, bitSize=64
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	//in this point we have a valid user_id
	user, apiErr := services.UsersService.GetUser(userID)
	if apiErr != nil {
		// Handle the err and return to the client
		utils.RespondError(c, apiErr)
		return
	}
	//by using JSON we set the header as application/json
	//if i´m located here, i know the userIDParam was valid and the user was found
	//so i can just return it
	utils.Respond(c, http.StatusOK, user)
}
