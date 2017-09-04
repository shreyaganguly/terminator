package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/go-fsnotify/fsnotify"
)

var logs []string

func watchLogDir(file *os.File, wordsList [][]string) {
	if file != nil {
		fmt.Println("Started watching your logs and will write all alerts to ", file.Name())
	} else {
		fmt.Println("Started watching your logs and will write all alerts to console")
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Terminator Error: ", err)
	}
	defer watcher.Close()

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				if event.Op.String() == "WRITE" {
					searchForKeyWords(file, wordsList, event.Name)
				}

				// watch for errors
			case err := <-watcher.Errors:
				log.Fatal("Terminator Error: ", err)
			}
		}
	}()

	if err := watcher.Add(tempDir); err != nil {
		log.Fatal("Terminator Error: ", err)
	}

	<-done

}

func createFilterFile() *os.File {
	f, err := os.Create(*filterFileName)
	if err != nil {
		log.Fatal("Terminator Error: ", err)
		return nil
	}
	return f
}

func notifyUser(file *os.File, line, filename string) {
	if len(*filterFileName) == 0 {
		fmt.Println("ALERT ALERT ALERT    ", line)

	} else {
		w := bufio.NewWriter(file)
		_, err := w.WriteString(line + "\n")
		if err != nil {
			log.Fatal("Terminator Error: ", err)
		}
		// Use `Flush` to ensure all buffered operations have
		// been applied to the underlying writer.
		w.Flush()
	}
	logs = append(logs, line)
	lr.Reload(filename)

}
