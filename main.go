package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

func main() {
	flag.Parse()
	var newWindow bool
	readCommands()
	for _, v := range commands {
		var i int
		var err error
		if newWindow == false {
			cmd := exec.Command("osascript", "-e", "tell application \"System Events\" to tell process \"Terminal\" to keystroke \"t\" using command down", "-e", fmt.Sprintf("tell application \"Terminal\" to do script \"%s\" in front window", v))
			var out bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err = cmd.Run()
			if err != nil {
				fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
				return
			}
			i, err = strconv.Atoi(strings.Split(out.String(), " ")[1])
			if err != nil {
				fmt.Println(err.Error())
				return
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
