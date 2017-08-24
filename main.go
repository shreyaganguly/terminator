package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	file = flag.String("f", "", "Filename where set of commands are given")
)

var (
	commands []string
)

func main() {
	flag.Parse()
	readCommands()
	for _, v := range commands {
		cmd := exec.Command("osascript", "-e", "tell application \"System Events\" to tell process \"Terminal\" to keystroke \"t\" using command down", "-e", fmt.Sprintf("tell application \"Terminal\" to do script \"%s\" in front window", v))
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return
		}
		fmt.Println("Result: " + out.String())
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
