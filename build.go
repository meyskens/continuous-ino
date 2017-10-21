package main

import (
	"errors"
	"os"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func clone(repo, hash string) (string, error) {
	nameParts := strings.Split(repo, "/")
	name := nameParts[len(nameParts)-1]
	path := "/tmp/" + name

	os.RemoveAll(path)
	r, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      "https://github.com/src-d/go-git",
		Progress: os.Stdout,
	})
	if err != nil {
		return "", err
	}

	w, err := r.Worktree()
	if err != nil {
		return "", err
	}

	err = w.Checkout(&git.CheckoutOptions{Hash: plumbing.NewHash(hash)})
	if err != nil {
		return "", err
	}

	return path, nil
}

func checkBuildFile(path string) error {
	_, err := os.Open(path + "/.cino.yml")
	if err != nil {
		return errors.New(".cino.yml not present")
	}

	return nil
}
