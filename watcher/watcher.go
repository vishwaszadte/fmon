package watcher

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/vishwaszadte/fmon/process"
)

// Watch starts the watcher and runs the initial command
// It restarts the command on file changes
func Watch(dir string, command string) {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Run the initial command
	process.Start(command)

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Printf("File modified: %s\n", event.Name)
				process.Restart(command)
			}
		case err := <-watcher.Errors:
			log.Println("Watcher error:", err)
		}
	}

}
