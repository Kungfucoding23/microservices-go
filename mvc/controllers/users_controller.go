package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Kungfucoding23/microservices-go/mvc/services"
	"github.com/Kungfucoding23/microservices-go/mvc/utils"
)

// GetUser controller
func GetUser(w http.ResponseWriter, r *http.Request) {
	userIDParam := r.URL.Query().Get("user_id")
	if len(userIDParam) < 1 {
		http.Error(w, "user id is needed", http.StatusBadRequest)
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
		jsonValue, _ := json.Marshal(apiErr)
		// Just return the Bad Request to the client.
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		return
	}
	user, apiErr := services.UsersService.GetUser(userID)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))
		// Handle the err and return to the client
		return
	}
	//return user to client
	//set header
	w.Header().Set("Content-Type", "application/json")
	//if i´m located here, i know the userIDParam was valid and the user was found
	//so i can just return it
	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)
}
