package main

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var gitHubClient *github.Client

func setUpGitHub() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	gitHubClient = github.NewClient(tc)
}
