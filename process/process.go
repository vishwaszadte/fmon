package process

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/vishwaszadte/fmon/utils"
)

var cmd *exec.Cmd

func Start(command string) {
	fmt.Printf("%s %s\n", utils.ForegroundColorPrimary("Starting process:"), command)
	cmd = exec.Command("bash", "-c", command)
	cmd.Stdout = nil
	cmd.Stderr = nil
	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to start process: %s\n", err)
	}
}

func Stop() {
	if cmd != nil && cmd.Process != nil {
		fmt.Println("Stopping process")
		cmd.Process.Kill()
		cmd.Wait()
	}
}

func Restart(command string) {
	Stop()
	time.Sleep(1 * time.Second)
	Start(command)
}
