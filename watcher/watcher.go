package watcher

import (
	"fmt"
	"log"
	"time"

	"github.com/radovskyb/watcher"
	"github.com/vishwaszadte/fmon/process"
	"github.com/vishwaszadte/fmon/utils"
)

type Watcher struct {
	Dir     string
	Command string
}

// Watch starts the watcher and runs the initial command
// It restarts the command on file changes
func (w *Watcher) Watch() {
	wtch := watcher.New()

	defer wtch.Close()

	wtch.SetMaxEvents(1)
	wtch.FilterOps(watcher.Write)

	go func() {
		for {
			select {
			case event := <-wtch.Event:
				// fmt.Printf("%s %s\n", utils.ForegroundColorPrimary("File modified:"), event.files[0])
				fmt.Printf("\n%s %s\n", utils.ForegroundColorPrimary("File modified:"), event.Path)
				process.Restart(w.Command)
			case err := <-wtch.Error:
				log.Println("Watcher error:", err)
			case <-wtch.Closed:

				fmt.Printf("%s\n", utils.ForegroundColorPrimary("Exiting..."))
			}
		}
	}()

	// Run the initial command
	process.Start(w.Command)

	if err := wtch.AddRecursive(w.Dir); err != nil {

		log.Fatal(err)
	}

	if err := wtch.Start(time.Millisecond * 100); err != nil {
		log.Fatal(err)
	}

}
