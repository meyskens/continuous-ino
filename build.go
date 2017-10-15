package main

import (
	"os"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
)

func clone(repo string) (string, error) {
	nameParts := strings.Split(repo, "/")
	name := nameParts[len(nameParts)-1]
	path := "/tmp/" + name

	os.RemoveAll(path)
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      "https://github.com/src-d/go-git",
		Progress: os.Stdout,
	})
	if err != nil {
		return "", err
	}

	return name, nil
}
