package src

import (
	"fmt"
	"os"
	"os/exec"
)

func checkService(name string) {
	cmd := exec.Command("systemctl", "status", name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("systemctl finished with non-zero: %v\n", exitErr)
		} else {
			fmt.Printf("failed to run systemctl: %v", err)
			os.Exit(1)
		}
	}
	fmt.Printf("Status is: %s\n", string(out))
}
