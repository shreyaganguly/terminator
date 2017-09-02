package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func commandExec(cmd *exec.Cmd) bytes.Buffer {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal("Terminator Error: ", fmt.Sprint(err)+": "+stderr.String())
		return out
	}
	return out
}

func readCommands() []string {
	var commands []string
	content, err := os.Open(*fileName)
	if err != nil {
		log.Fatal("Terminator Error: ", err.Error())
		return commands
	}
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}
	return commands
}
