package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "SECRET_GITHUB_ACCESS_TOKEN", githubAccessTokenEnvName)
}

func TestGetGithubAccessToken(t *testing.T) {
	assert.EqualValues(t, "", GetGithubAccessToken())
}
