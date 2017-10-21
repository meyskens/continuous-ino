package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var gitHubClient *github.Client
var buildMutex = sync.Mutex{}

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
	setPending(payload)

	buildMutex.Lock()

	fs, err := clone(payload.Repo.GetCloneURL(), payload.GetAfter())
	if err != nil {
		fmt.Println(err)
		setFailed(payload, "Could not clone repository")
		return
	}

	err = checkBuildFile(fs)
	if err != nil {
		setFailed(payload, err.Error())
		return
	}

	buildMutex.Unlock()

	setSuccess(payload)
}

func setPending(payload github.WebHookPayload) {
	gitHubClient.Repositories.CreateStatus(context.Background(), payload.Repo.Owner.GetLogin(), payload.Repo.GetName(), payload.GetAfter(), &github.RepoStatus{State: &statusPending})
}

func setFailed(payload github.WebHookPayload, reason string) {
	gitHubClient.Repositories.CreateStatus(context.Background(), payload.Repo.Owner.GetLogin(), payload.Repo.GetName(), payload.GetAfter(), &github.RepoStatus{State: &statusFailure, Context: &reason})
}

func setSuccess(payload github.WebHookPayload) {
	gitHubClient.Repositories.CreateStatus(context.Background(), payload.Repo.Owner.GetLogin(), payload.Repo.GetName(), payload.GetAfter(), &github.RepoStatus{State: &statusSuccess})
}
