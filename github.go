package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var gitHubClient *github.Client

var statusPending = "pending"
var statusSuccess = "success"
var statusError = "error"
var statusFailure = "failure"

func setUpGitHub() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GitHub.AuthToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	gitHubClient = github.NewClient(tc)
}

func handleHookPayload(event string, payload github.WebHookPayload) {
	fmt.Println("got event", event)
	switch event {
	case "push":
		go handlePush(payload)
	default:
		return
	}
}

func handlePush(payload github.WebHookPayload) {
	// Set pending
	_, _, err := gitHubClient.Repositories.CreateStatus(context.Background(), payload.Repo.Owner.GetLogin(), payload.Repo.GetName(), payload.GetRef(), &github.RepoStatus{State: &statusPending})
	fmt.Println(err)

	// dummy it for now
	time.Sleep(10 * time.Second)
	_, _, err = gitHubClient.Repositories.CreateStatus(context.Background(), payload.Repo.Owner.GetLogin(), payload.Repo.GetName(), payload.GetRef(), &github.RepoStatus{State: &statusSuccess})
	fmt.Println(err)
}
