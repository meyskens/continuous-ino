package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/src-d/go-billy.v3"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func clone(repo, hash string) (billy.Filesystem, error) {
	nameParts := strings.Split(repo, "/")
	name := nameParts[len(nameParts)-1]
	path := "/tmp/" + name

	os.RemoveAll(path)
	r, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      repo,
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println("clone error", err)
		return nil, err
	}

	w, err := r.Worktree()
	if err != nil {
		fmt.Println("worktree error", err)
		return nil, err
	}

	err = w.Checkout(&git.CheckoutOptions{Force: true, Hash: plumbing.NewHash(hash)})
	if err != nil {
		fmt.Println("checkout error", err)
		return nil, err
	}

	return w.Filesystem, nil
}

func checkBuildFile(fs billy.Filesystem) error {
	_, err := fs.Open(".cino.yml")
	if err != nil {
		return errors.New(".cino.yml not present")
	}

	return nil
}
