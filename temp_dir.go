package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var tempDir string

func createTempDir() {
	var err error
	c, err := os.Getwd()
	if err != nil {
		log.Fatal("Terminator Error: ", err.Error())
	}
	tempDir, err = ioutil.TempDir(c, "_temp")
	if err != nil {
		log.Fatal("Terminator Error: ", err.Error())
	}
}

func createTempFile() *os.File {
	file, err := ioutil.TempFile(tempDir, "temp")
	if err != nil {
		log.Fatal("Terminator Error: ", err.Error())
	}
	return file
}

func cleanUp() {
	for _, v := range windowIDs {
		cmd := exec.Command("osascript", "-e", "tell application \"Terminal\"", "-e", fmt.Sprintf("close (every window whose id is %s)", v), "-e", "end tell")
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			log.Fatal("Terminator Error: ", fmt.Sprint(err)+": "+stderr.String())
		}
	}

	os.RemoveAll(tempDir)
}
