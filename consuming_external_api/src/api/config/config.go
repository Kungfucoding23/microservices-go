package config

import "os"

const (
	apiGithubToken = "SECRET_GITHUB_TOKEN"
)

var (
	githubToken = os.Getenv(apiGithubToken)
)

//GetGithubToken returns the github access token
func GetGithubToken() string {
	return githubToken
}
