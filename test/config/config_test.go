package configtest

import (
	"app/cmd/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	mode := os.Getenv("GIN_MODE")
	configfolder := os.Getenv("config")
	configs := config.Load(mode, configfolder)
	assert.Equal(t, "43.156.36.211", configs.Db.Host)
	assert.Equal(t, "auth_token", configs.Auth.Token)
	assert.Equal(t, "https://example.com/jira", configs.URL.Jira)
}
