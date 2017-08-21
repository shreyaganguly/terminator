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
	fmt.Println(len(commands))
	for _, v := range commands {
		cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Terminal\" to do script \"%s\"", v))
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
