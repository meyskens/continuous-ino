package config

import (
	"os"
)

// GitHubConfig contains the config specific to GitHub
type GitHubConfig struct {
	AuthToken     string
	WebhookSecret string
}

// ArduinoConfig is the configuration for used in uploading the Arduino code
type ArduinoConfig struct {
	Model string
}

// DatabaseConfig is the configuration used for BoltDB
type DatabaseConfig struct {
	Path string
}

// Configuration contains the config for the CI
type Configuration struct {
	GitHub   GitHubConfig
	Arduino  ArduinoConfig
	Database DatabaseConfig
}

// GetConfiguration reads the configuration from config.json and returns it
func GetConfiguration() Configuration {
	returnConfig := Configuration{
		Arduino: ArduinoConfig{
			Model: "nano", // set default model
		},
		Database: DatabaseConfig{
			Path: "/var/lib/cino/cino.db",
		},
	}

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
	if model := os.Getenv("CINO_ARDUINO_MODEL"); model != "" {
		conf.Arduino.Model = model
	}
	if path := os.Getenv("CINO_DATABASE_PATH"); path != "" {
		conf.Database.Path = path
	}
}
