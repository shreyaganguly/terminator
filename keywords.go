package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var filteringWords [][]string

func trimSpace(w []string) []string {
	var trimmedWords []string
	for _, v := range w {
		trimmedWords = append(trimmedWords, strings.TrimSpace(v))
	}
	return trimmedWords
}

func getWords() [][]string {
	var wordsList [][]string
	words := strings.Split(*keyWords, ",")
	fmt.Println("Following are your filtering keywords")
	for _, v := range words {
		andWords := make([]string, len(trimSpace(strings.Split(v, "and"))))
		copy(andWords, trimSpace(strings.Split(v, "and")))
		wordsList = append(wordsList, andWords)
		fmt.Println(andWords)
	}
	return wordsList
}

func searchForKeyWords(filename string) {
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
