package buildfile

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	testFile, _ := ioutil.ReadFile("test.yml")

	type args struct {
		content []byte
	}
	tests := []struct {
		name    string
		args    args
		want    BuildFile
		wantErr bool
	}{
		{
			name: "success test",
			args: args{
				content: testFile,
			},
			want: BuildFile{
				Main: "test.ino",
				Tests: []TestContent{
					{Name: "test1"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
