package main

import (
	"io"
	"os"
	"os/exec"

	"github.com/meyskens/continuous-ino/buildfile"
)

func runTests(path string, buildFile buildfile.BuildFile) (succeeded bool, errs []error) {
	errs = []error{}
	succeeded = true

	for _, test := range buildFile.Tests {
		err := buildAndTestIno(path, buildFile, test)
		if err != nil {
			succeeded = false
			errs = append(errs, err)
		}
	}

	return
}

func buildAndTestIno(path string, buildFile buildfile.BuildFile, test buildfile.TestContent) (err error) {
	path = path + "/"
	// Backup main.ino
	os.Rename(path+buildFile.Main, path+buildFile.Main+".bak")
	// Copy over test file
	copyFile(path+test.File, path+buildFile.Main)

	cmd := exec.Command("/bin/bash", "-c", "cd "+path+" && ino build -m "+cfg.Arduino.Model)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err == nil {
		// No errors! we can run!
		cmd := exec.Command("/bin/bash", "-c", "cd "+path+" && ino upload -m "+cfg.Arduino.Model)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}

	// Remove test file
	os.Remove(path + buildFile.Main)
	// Restore main.ino
	os.Rename(path+buildFile.Main+".bak", path+buildFile.Main)

	return
}

func copyFile(src, dst string) {
	from, _ := os.Open(src)
	defer from.Close()

	to, _ := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0755)
	defer to.Close()

	io.Copy(to, from)
}
