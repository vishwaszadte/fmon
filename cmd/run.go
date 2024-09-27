package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vishwaszadte/fmon/watcher"
)

var runCmd = &cobra.Command{
	Use:   "run [command]",
	Short: "Run a command and re-run it on file changes",
	Args:  cobra.MinimumNArgs(1),
	Run:   cmdHelper,
}

func cmdHelper(cmd *cobra.Command, args []string) {
	fmt.Println("Monitoring for file changes...")
	command := args[0]

	// Start watching files
	go watcher.Watch(".", command)

	// Block the main thread until terminated
	select {}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
