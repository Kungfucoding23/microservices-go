package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "SECRET_GITHUB_TOKEN", apiGithubToken)
}

func TestGetGithubTokenNoError(t *testing.T) {
	result := GetGithubToken()
	assert.NotNil(t, result)
	assert.EqualValues(t, "", result)
}
