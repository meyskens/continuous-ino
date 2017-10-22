package config

import (
	"os"
	"reflect"
	"testing"
)

func Test_readEnv(t *testing.T) {
	type args struct {
		conf *Configuration
	}
	tests := []struct {
		name    string
		args    args
		envName string
		envVal  string
		want    *Configuration
	}{
		{
			name:    "test github token",
			args:    args{conf: &Configuration{}},
			envName: "CINO_GITHUB_AUTHTOKEN",
			envVal:  "abc123",
			want:    &Configuration{GitHub: GitHubConfig{AuthToken: "abc123"}},
		},
		{
			name:    "test github token",
			args:    args{conf: &Configuration{}},
			envName: "CINO_GITHUB_WEBHOOK_SECRET",
			envVal:  "secreeet",
			want:    &Configuration{GitHub: GitHubConfig{WebhookSecret: "secreeet"}},
		},
		{
			name:    "test arduino model",
			args:    args{conf: &Configuration{}},
			envName: "CINO_ARDUINO_MODEL",
			envVal:  "uno",
			want:    &Configuration{Arduino: ArduinoConfig{Model: "uno"}},
		},
		{
			name:    "test database path",
			args:    args{conf: &Configuration{}},
			envName: "CINO_DATABASE_PATH",
			envVal:  "/tmp/db",
			want:    &Configuration{Database: DatabaseConfig{Path: "/tmp/db"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(tt.envName, tt.envVal)
			if readEnv(tt.args.conf); !reflect.DeepEqual(tt.args.conf, tt.want) {
				t.Errorf("readEnv() = %v, want %v", tt.args.conf, tt.want)
			}
			os.Unsetenv(tt.envName)
		})
	}
}
