package main

import (
	"reflect"
	"testing"

	"github.com/meyskens/continuous-ino/buildfile"
	"gopkg.in/src-d/go-billy.v3"
	"gopkg.in/src-d/go-billy.v3/memfs"
)

func Test_clone(t *testing.T) {
	type args struct {
		repo string
		hash string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test checkout",
			args: args{
				repo: "https://github.com/meyskens/cino-test.git",
				hash: "1432c32e4b9e900befd9edb4982fad2f000496ce",
			},
			wantErr: false,
		},
		{
			name: "test fake hash",
			args: args{
				repo: "https://github.com/meyskens/cino-test.git",
				hash: "1432c32e4b9e900befd9edb4982fad2f00049dd6ce",
			},
			wantErr: true,
		},
		{
			name: "test fake repo",
			args: args{
				repo: "https://github.com/meyskens/cino-fake.git",
				hash: "1432c32e4b9e900befd9edb4982fad2f000496ce",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := clone(tt.args.repo, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("clone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_readBuildFile(t *testing.T) {
	fs1 := memfs.New()
	fs2 := memfs.New()
	fs1.Create(".cino.yml")

	type args struct {
		fs billy.Filesystem
	}
	tests := []struct {
		name    string
		args    args
		wantOut buildfile.BuildFile
		wantErr bool
	}{
		{
			name: "test positive",
			args: args{
				fs: fs1,
			},
			wantErr: false,
		},
		{
			name: "test negative",
			args: args{
				fs: fs2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := readBuildFile(tt.args.fs)
			if (err != nil) != tt.wantErr {
				t.Errorf("readBuildFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("readBuildFile() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
