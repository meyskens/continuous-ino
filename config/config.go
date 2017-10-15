package config

import (
	"os"
)

// GitHubConfig contains the config specific to GitHub
type GitHubConfig struct {
	AuthToken     string
	WebhookSecret string
}

// Configuration contains the config for the CI
type Configuration struct {
	GitHub GitHubConfig
}

// GetConfiguration reads the configuration from config.json and returns it
func GetConfiguration() Configuration {
	returnConfig := Configuration{}

	readEnv(&returnConfig)
	return returnConfig
}

func readEnv(conf *Configuration) {
	if authtoken := os.Getenv("CINO_GITHUB_AUTHTOKEN"); authtoken != "" {
		conf.GitHub.AuthToken = authtoken
	}
	if secret := os.Getenv("CINO_GITHUB_WEBHOOK_SECRET"); secret != "" {
		conf.GitHub.WebhookSecret = secret
	}
}
