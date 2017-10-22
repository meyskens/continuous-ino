package main

import (
	"os"
	"testing"

	"github.com/meyskens/continuous-ino/buildfile"
	"github.com/meyskens/continuous-ino/config"
)

const TestPath = "/tmp/cino-test/"
const TestFile = `
void setup() {
	// put your setup code here, to run once:
  
}
  
void loop() {
	// put your main code here, to run repeatedly:
  
}
`

const TestFileBad = `
void setup() {
	// put your setup code here, to run once:
  
}
  
void loop() {
	// put your main code here, to run repeatedly:
  
`

func init() {
	cfg = config.GetConfiguration()
}

var TestBuildFile = buildfile.BuildFile{
	Main: "./src/main.ino",
	Tests: []buildfile.TestContent{
		{
			Name: "test 1",
			File: "./test/test.ino",
		},
		{
			Name: "test bad",
			File: "./test/test-bad.ino",
		},
	},
}

func Test_buildAndTestIno(t *testing.T) {
	os.RemoveAll(TestPath)
	os.MkdirAll(TestPath, 0755)
	os.MkdirAll(TestPath+"/test/", 0755)
	os.MkdirAll(TestPath+"/src/", 0755)

	f, _ := os.Create(TestPath + TestBuildFile.Main)
	f.WriteString(TestFile)
	f.Close()

	f, _ = os.Create(TestPath + TestBuildFile.Tests[0].File)
	f.WriteString(TestFile)
	f.Close()

	f, _ = os.Create(TestPath + TestBuildFile.Tests[1].File)
	f.WriteString(TestFileBad)
	f.Close()

	type args struct {
		path      string
		buildFile buildfile.BuildFile
		test      buildfile.TestContent
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test build",
			args: args{
				path:      TestPath,
				buildFile: TestBuildFile,
				test:      TestBuildFile.Tests[0],
			},
			wantErr: false,
		},
		{
			name: "test fail build",
			args: args{
				path:      TestPath,
				buildFile: TestBuildFile,
				test:      TestBuildFile.Tests[1],
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := buildAndTestIno(tt.args.path, tt.args.buildFile, tt.args.test); (err != nil) != tt.wantErr {
				t.Errorf("buildAndTestIno() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	os.RemoveAll(TestPath)
}
