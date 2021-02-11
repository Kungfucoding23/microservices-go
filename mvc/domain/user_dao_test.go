package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	//Steps:
	//Initialization: in this case I dont need to initialize anything

	//Execution
	user, err := GetUser(0)

	//Validation
	//This line replaces the if sentence
	assert.Nil(t, user, "We were not expecting a user with id 0")

	// if user != nil {
	// 	t.Error("We were not expecting a user with id 0")
	// }

	assert.NotNil(t, err, "We were expecting an error when user id is 0")

	// if err == nil {
	// 	t.Error("We were expecting an error when user id is 0")
	// }
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	// if err.StatusCode != http.StatusNotFound {
	// 	t.Error("We were expecting 404 when user is not found")
	// }
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User 0 does not exists", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	//I´m not expecting any error
	assert.Nil(t, err)
	//and a not nill user
	assert.NotNil(t, user)
	//I expect the user 123
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Ale", user.FirstName)
	assert.EqualValues(t, "Ale", user.FirstName)
	assert.EqualValues(t, "Santin", user.LastName)
	assert.EqualValues(t, "alejandro@gmail.com", user.Email)
}
