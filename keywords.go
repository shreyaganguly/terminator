package main

import (
	"bufio"
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

func searchForKeyWords(logFile *os.File, wordsList [][]string, filename string) {
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
				filterText(logFile, wordsList, scanner.Text(), filename)
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

func filterText(file *os.File, wordsList [][]string, line string, filename string) {
	for _, orWords := range wordsList {
		var found = 1
		for _, andWord := range orWords {
			if !strings.Contains(line, andWord) {
				found = 0
			}
		}
		if found == 1 {
			notifyUser(file, line, filename)
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
