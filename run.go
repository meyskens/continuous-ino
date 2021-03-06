package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/jacobsa/go-serial/serial"
	"github.com/meyskens/continuous-ino/serialhandler"
	"github.com/meyskens/continuous-ino/storage"

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
			// save to db in real time bro
			currentRun.Errors = errs
			store.SaveRun(currentRun)
		}
	}

	currentRun.Running = false
	store.SaveRun(currentRun)
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

	runOutput := storage.RunOutput{}
	runOutput.File = test.File
	runOutput.Name = test.Name
	runOutput.Step = "Upload project"

	_, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	handler := serialhandler.New(cancel)

	// No errors! we can upload!
	cmd := execCommand("/bin/bash", "-c", "cd "+path+" && arduino --upload --pref sketchbook.path=$(pwd) --board "+cfg.Arduino.Model+" --port "+cfg.Arduino.Port+" "+test.File)
	cmd.Stdout = &handler
	cmd.Stderr = &handler
	err = cmd.Run()

	fmt.Println(handler.Output())
	runOutput.Output += handler.Output()

	currentRun.Output = append(currentRun.Output, runOutput)
	store.SaveRun(currentRun)

	if err != nil {
		fmt.Println(err)
		return
	}

	// No errors! we can run!
	runOutput = storage.RunOutput{}
	runOutput.File = test.File
	runOutput.Name = test.Name
	runOutput.Step = "Run on Arduino"

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	handler = serialhandler.New(cancel)

	port, err := serial.Open(serial.OpenOptions{
		PortName:        cfg.Arduino.Port,
		BaudRate:        uint(buildFile.Baud),
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	})

	if err != nil {
		return
	}

	defer port.Close()
	go pipe(port, &handler)

	<-ctx.Done()
	fmt.Println(handler.Output())
	runOutput.Output += handler.Output()

	if len(handler.Errors()) != 0 {
		err = errors.New(strings.Join(handler.Errors(), "\n"))
		fmt.Println("Found errors")
		fmt.Println(err)
	}

	currentRun.Output = append(currentRun.Output, runOutput)
	store.SaveRun(currentRun)

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
