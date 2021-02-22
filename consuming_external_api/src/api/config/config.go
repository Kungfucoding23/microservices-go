package config

import "os"

const (
	apiGithubToken = "SECRET_GITHUB_TOKEN"
	//LogLevel ..
	LogLevel      = "info"
	goEnvironment = "GO_ENVIRONMENT"
	production    = "production"
)

var (
	githubToken = os.Getenv(apiGithubToken)
)

//GetGithubToken returns the github access token
func GetGithubToken() string {
	return githubToken
}

//IsProduction ..
func IsProduction() bool {
	return os.Getenv(goEnvironment) == production
}
