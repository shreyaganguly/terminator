package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func searchForKeyWord(filename string) {
	var lastLine string
	if file, err := os.Open(filename); err == nil {

		defer file.Close()

		scanner := bufio.NewScanner(file)
		found := 0
		for scanner.Scan() {
			lastLine = scanner.Text()
			if found == 0 {
				if fileMap[filename] == scanner.Text() {
					found = 1
				}
			}
			if found == 1 {
				if strings.Contains(scanner.Text(), "card") {
					fmt.Println("ALERT ALERT ALERT", scanner.Text())
				}
			}
		}
		fileMap[filename] = lastLine

		if err = scanner.Err(); err != nil {
			log.Fatal("Terminator Error: ", err.Error())
		}

	} else {
		log.Fatal("Terminator Error: ", err.Error())
	}
}
