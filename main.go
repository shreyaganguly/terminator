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
	file = flag.String("f", "", "Filename where set of commands are given")
)

var (
	commands []string
)

const (
	tabLimit = 10
)

func createTempFile() *os.File {
	tempDir := os.TempDir()
	// c, _ := os.Getwd()
	file, err := ioutil.TempFile(tempDir, "temp")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Temp File created!")
	// defer os.Remove(file.Name())
	return file
}

func main() {
	flag.Parse()
	var newWindow bool
	readCommands()
	file := createTempFile()
	for _, v := range commands {
		var i int
		var err error

		if newWindow == false {
			cmd := exec.Command("osascript", "-e", "tell application \"System Events\" to tell process \"Terminal\" to keystroke \"t\" using command down", "-e", fmt.Sprintf("tell application \"Terminal\" to do script \"%s | tee %s\" in front window", v, file.Name()))
			var out bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err = cmd.Run()
			if err != nil {
				fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			}
			_, err = strconv.Atoi(strings.Split(out.String(), " ")[1])
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Terminal\" to do script \"%s\"", v))
			var out bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err = cmd.Run()
			if err != nil {
				fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
				return
			}
			newWindow = false

		}

		if i > tabLimit {
			newWindow = true
		}
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		os.Remove(file.Name())
		os.Exit(0)
	}()
	http.ListenAndServe(":8080", nil)

}

func readCommands() {
	content, err := os.Open(*file)
	if err != nil {
		log.Fatal("Error")
	}
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
}
