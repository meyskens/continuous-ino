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
