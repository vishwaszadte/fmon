package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/vishwaszadte/fmon/utils"
	"github.com/vishwaszadte/fmon/watcher"
)

// includeDir is the variable that stores the directory to watch
var includeDir string

var runCmd = &cobra.Command{
	Use:   "run [command]",
	Short: "Run a command and re-run it on file changes",
	Args:  cobra.MinimumNArgs(1),
	Run:   runCmdHelper,
}

// runCmdHelper is the main function that is called when the run command is executed
// It starts the watcher and blocks the main thread until terminated
func runCmdHelper(cmd *cobra.Command, args []string) {

	fmt.Printf("%s\n", utils.ForegroundColorPrimary("Monitoring for file changes..."))

	absDir, err := filepath.Abs(includeDir)
	if err != nil {
		log.Fatalf("Invalid directory: %s", err.Error())
	}

	fmt.Printf("%s %s\n", utils.ForegroundColorPrimary("Watching directory:"), absDir)
	command := args[0]

	// Start watching files
	go watcher.Watch(absDir, command)

	// Block the main thread until terminated
	select {}
}

func init() {

	// Define the --incl-dir flag, which is used to specify the directory to watch
	runCmd.Flags().StringVarP(&includeDir, "incl-dir", "i", ".", "Directory to watch (default is current directory)")

	rootCmd.AddCommand(runCmd)
}
