package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/jacobsa/go-serial/serial"
	"github.com/meyskens/continuous-ino/serialhandler"

	"github.com/meyskens/continuous-ino/buildfile"
)

var execCommand = exec.Command
var execCommandContext = exec.CommandContext

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
	var timeout time.Duration
	if test.Timeout != "" {
		timeout, err = time.ParseDuration(test.Timeout)
		if err != nil {
			return
		}
	} else {
		timeout = time.Minute * 10
	}
	path = path + "/"
	// Backup main.ino
	os.Rename(path+buildFile.Main, path+buildFile.Main+".bak")
	// Copy over test file
	copyFile(path+test.File, path+buildFile.Main)

	cmd := execCommand("/bin/bash", "-c", "cd "+path+" && ino build -m "+cfg.Arduino.Model)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err == nil {
		// No errors! we can upload!
		cmd = execCommand("/bin/bash", "-c", "cd "+path+" && ino upload -m "+cfg.Arduino.Model)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
	}

	if err == nil {
		// No errors! we can run!
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		handler := serialhandler.New(cancel)

		port, err := serial.Open(serial.OpenOptions{
			PortName:        "/dev/ttyUSB0",
			BaudRate:        uint(buildFile.Baud),
			DataBits:        8,
			StopBits:        1,
			MinimumReadSize: 4,
		})

		if err == nil {
			defer port.Close()
			go pipe(port, &handler)

			<-ctx.Done()
			fmt.Println(handler.Output())
		}
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

func pipe(r io.Reader, w io.Writer) {
	var err error
	var n int
	for err == nil {
		p := make([]byte, 1)
		n, err = r.Read(p)
		w.Write(p[0:n])
	}
}
