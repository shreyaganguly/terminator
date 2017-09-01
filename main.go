package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

var (
	fileName = flag.String("f", "", "Filename where set of commands are given")
	keyWords = flag.String("words", "", "Keywords to search for in the logs being monitored(separated with \",\")")
)

var (
	commands  []string
	tempDir   string
	windowIDs []string
	fileMap   map[string]int
)

const (
	tabLimit = 10
)

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

func main() {
	flag.Parse()
	var newWindow bool
	filteringWords = getWords()
	fileMap = make(map[string]int)
	createTempDir()

	var file *os.File
	var notFirstWindow bool
	readCommands()
	var i int
	for _, v := range commands {

		var err error
		file = createTempFile()
		if i > (tabLimit - 1) {
			newWindow = true
		}

		if newWindow == false && notFirstWindow {
			cmd := exec.Command("osascript", "-e", "tell application \"System Events\" to tell process \"Terminal\" to keystroke \"t\" using command down", "-e", fmt.Sprintf("tell application \"Terminal\" to do script \"%s | tee %s\" in front window", v, file.Name()))
			var out bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err = cmd.Run()
			if err != nil {
				log.Fatal("Terminator Error: ", fmt.Sprint(err)+": "+stderr.String())
			}
			i, err = strconv.Atoi(strings.Split(out.String(), " ")[1])
			if err != nil {
				log.Fatal("Terminator Error: ", err.Error())
			}
		} else {
			cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Terminal\" to do script \"%s| tee %s\"", v, file.Name()))
			var out bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err = cmd.Run()
			if err != nil {
				log.Fatal("Terminator Error: ", fmt.Sprint(err)+": "+stderr.String())
			}
			windowID := strings.TrimSpace(strings.Split(out.String(), " ")[5])
			i, err = strconv.Atoi(strings.Split(out.String(), " ")[1])
			if err != nil {
				log.Fatal("Terminator Error: ", err.Error())
			}
			windowIDs = append(windowIDs, windowID)
			newWindow = false
			notFirstWindow = true
		}

	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		cleanUp()
		os.Exit(0)
	}()
	http.ListenAndServe(":8080", nil)

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

func readCommands() {
	content, err := os.Open(*fileName)
	if err != nil {
		log.Fatal("Terminator Error: ", err.Error())
	}
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}
}
