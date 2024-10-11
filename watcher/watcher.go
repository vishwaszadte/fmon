package watcher

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/vishwaszadte/fmon/process"
	"github.com/vishwaszadte/fmon/utils"
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
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Has(fsnotify.Write) {
				fmt.Printf("%s %s\n", utils.ForegroundColorPrimary("File modified:"), event.Name)
				process.Restart(command)
			}
		case err := <-watcher.Errors:
			log.Println("Watcher error:", err)
		}
	}

}
