package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/go-github/github"
	"github.com/meyskens/continuous-ino/storage"
	"golang.org/x/oauth2"
)

var gitHubClient *github.Client
var buildMutex = sync.Mutex{}

var currentRun storage.Run // we only run one at the time for now!

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
	currentRun = store.NewRun()
	currentRun.Repo = payload.Repo.GetFullName()
	currentRun.Time = time.Now()
	currentRun.Output = []storage.RunOutput{}
	store.SaveRun(currentRun)

	setPending(payload)

	buildMutex.Lock()

	fs, path, err := clone(payload.Repo.GetCloneURL(), payload.GetAfter())
	if err != nil {
		fmt.Println(err)
		setFailed(payload, "Could not clone repository")
		return
	}

	bfile, err := readBuildFile(fs)
	if err != nil {
		setFailed(payload, err.Error())
		return
	}

	ok, _ := runTests(path, bfile)
	if !ok {
		buildMutex.Unlock()
		setFailed(payload, "Some tests did not pass")
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
