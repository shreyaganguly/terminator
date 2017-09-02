package main

import (
	"flag"
	"fmt"
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
	host     = flag.String("b", "0.0.0.0", "Host to start the Terminator application")
	port     = flag.String("p", "8080", "Port to start the Terminator application")
)

var (
	windowIDs                 []string
	newWindow, notFirstWindow bool
	i                         int
)

const (
	tabLimit = 10
)

func main() {
	flag.Parse()
	filteringWords := getFilterKeyWords()
	commands := readCommands()
	createTempDir()
	if *keyWords != "" {
		go watchLogDir(filteringWords)
	}
	for _, v := range commands {
		var err error
		file := createTempFile()
		if i > (tabLimit - 1) {
			newWindow = true
		}
		if newWindow == false && notFirstWindow {
			result := commandExec(exec.Command("osascript", "-e", "tell application \"System Events\" to tell process \"Terminal\" to keystroke \"t\" using command down", "-e", fmt.Sprintf("tell application \"Terminal\" to do script \"%s | tee %s\" in front window", v, file.Name())))
			i, err = strconv.Atoi(strings.Split(result.String(), " ")[1])
			if err != nil {
				log.Fatal("Terminator Error: ", err.Error())
			}
		} else {
			result := commandExec(exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Terminal\" to do script \"%s| tee %s\"", v, file.Name())))
			windowID := strings.TrimSpace(strings.Split(result.String(), " ")[5])
			i, err = strconv.Atoi(strings.Split(result.String(), " ")[1])
			if err != nil {
				log.Fatal("Terminator Error: ", err.Error())
			}
			windowIDs = append(windowIDs, windowID)
			newWindow = false
			notFirstWindow = true
		}
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-sigs
		cleanUp()
		os.Exit(0)
	}()
	addr := fmt.Sprintf("%s:%s", *host, *port)
	fmt.Println("Starting the Terminator Application at", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Terminator Error: ", err.Error())
	}

}
