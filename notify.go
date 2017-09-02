package main

import (
	"fmt"
	"log"

	"github.com/go-fsnotify/fsnotify"
)

func watchLogDir(wordsList [][]string) {
	fmt.Println("Started watching your logs")
	// creates a new file watcher
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
					searchForKeyWords(wordsList, event.Name)
				}

				// watch for errors
			case err := <-watcher.Errors:
				log.Fatal("Terminator Error: ", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add(tempDir); err != nil {
		log.Fatal("Terminator Error: ", err)
	}

	<-done

}
