package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJSON(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "Intro to consuming external api",
		Description: "an intro to github api using golang",
		Homepage:    "https://github.com",
		Private:     false,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	//create a json
	//Marshal takes an input interface and attemps to create a valid json string
	bytes, err := json.Marshal(request)
	//not expecting an err
	assert.Nil(t, err)
	//expecting bytes
	assert.NotNil(t, bytes)

	var target CreateRepoRequest
	// Unmarshal takes an input byte array and a *pointer* that we´re trying to fill using this json
	err = json.Unmarshal(bytes, &target)
	// and we shouldn´t have any err because we are Marshal and Unmarshal the same data struct
	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)

	assert.EqualValues(t, `{"name":"Intro to consuming external api","description":"an intro to github api using golang","homepage":"https://github.com","private":false,"has_issues":true,"has_projects":true,"has_wiki":true}`, string(bytes))

}
