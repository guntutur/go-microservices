package config

import "os"

const (
	githubAccessTokenEnvName = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(githubAccessTokenEnvName)
)

func GetGithubAccessToken() string {
	return githubAccessToken
}
