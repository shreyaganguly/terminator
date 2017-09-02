package main

import (
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
		commandExec(exec.Command("osascript", "-e", "tell application \"Terminal\"", "-e", fmt.Sprintf("close (every window whose id is %s)", v), "-e", "end tell"))
	}

	os.RemoveAll(tempDir)
}
