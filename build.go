package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/meyskens/continuous-ino/buildfile"

	"gopkg.in/src-d/go-billy.v3"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func clone(repo, hash string) (fs billy.Filesystem, path string, err error) {
	nameParts := strings.Split(repo, "/")
	name := nameParts[len(nameParts)-1]
	path = "/tmp/" + name

	os.RemoveAll(path)
	r, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:               repo,
		Progress:          os.Stdout,
		Depth:             1,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
		fmt.Println("clone error", err)
		return
	}

	w, err := r.Worktree()
	if err != nil {
		fmt.Println("worktree error", err)
		return
	}

	err = w.Checkout(&git.CheckoutOptions{Force: true, Hash: plumbing.NewHash(hash)})
	if err != nil {
		fmt.Println("checkout error", err)
		return
	}

	fs = w.Filesystem

	return
}

func readBuildFile(fs billy.Filesystem) (out buildfile.BuildFile, err error) {
	file, err := fs.Open(".cino.yml")
	if err != nil {
		return
	}

	b, _ := ioutil.ReadAll(file)
	out, err = buildfile.Parse(b)

	return
}
