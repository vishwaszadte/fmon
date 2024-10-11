package process

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/vishwaszadte/fmon/utils"
)

var cmd *exec.Cmd

func Start(command string) {
	fmt.Printf("%s %s\n", utils.ForegroundColorPrimary("Starting process:"), command)
	cmd = exec.Command("bash", "-c", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to start process: %s\n", err)
	}

	fmt.Printf("%s: <%d>\n", utils.ForegroundColorPrimary("Process started"), cmd.Process.Pid)
}

func Stop() {
	if cmd == nil || cmd.Process == nil {
		fmt.Println("Process is not running")
	}

	fmt.Println(utils.ForegroundColorPrimary("Stopping process"))
	if err := cmd.Process.Kill(); err != nil {
		fmt.Printf("Failed to stop process: %s\n", err)
	}
	cmd.Wait()

}

func Restart(command string) {
	Stop()
	time.Sleep(1 * time.Second)
	Start(command)
}
