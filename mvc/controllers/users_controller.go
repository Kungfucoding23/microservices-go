package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Kungfucoding23/microservices-go/mvc/services"
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
		// Just return the Bad Request to the client.
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user_id must be a number"))
		return
	}
	user, err := services.GetUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		// Handle the err and return to the client
		return
	}
	//return user to client
	//set content type
	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)
}
