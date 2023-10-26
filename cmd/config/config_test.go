package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	configs := Load("devv")
	assert.Equal(t, "localhost", configs.Db.Host)
	assert.Equal(t, "auth_token", configs.Auth.Token)
	assert.Equal(t, "https://example.com/jira", configs.URL.Jira)
}
