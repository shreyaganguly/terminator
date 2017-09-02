package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	fileMap map[string]int
)

func getFilterKeyWords() [][]string {
	var wordsList [][]string
	fileMap = make(map[string]int)
	words := strings.Split(*keyWords, ",")
	for _, v := range words {
		andWords := make([]string, len(trimSpace(strings.Split(v, "and"))))
		copy(andWords, trimSpace(strings.Split(v, "and")))
		wordsList = append(wordsList, andWords)
	}
	return wordsList
}

func searchForKeyWords(wordsList [][]string, filename string) {
	var lastLine int
	if file, err := os.Open(filename); err == nil {

		defer file.Close()

		scanner := bufio.NewScanner(file)
		found := 0
		for scanner.Scan() {
			lastLine++
			if found == 0 {
				if fileMap[filename] < lastLine {
					found = 1
				}
			}
			if found == 1 {
				filterText(wordsList, scanner.Text())
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

func filterText(wordsList [][]string, line string) {
	for _, orWords := range wordsList {
		var found = 1
		for _, andWord := range orWords {
			if !strings.Contains(line, andWord) {
				found = 0
			}
		}
		if found == 1 {
			fmt.Println("ALERT ALERT ALERT    ", line)
		}
	}

}

func trimSpace(w []string) []string {
	var trimmedWords []string
	for _, v := range w {
		trimmedWords = append(trimmedWords, strings.TrimSpace(v))
	}
	return trimmedWords
}
