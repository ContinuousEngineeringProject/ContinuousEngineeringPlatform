
package main

import (
	"fmt"
	"os"
	"bufio"
	"os/exec"
)

func main() {

	// docker machine command
	cmdName := "docker-machine"
	cmdArgs := []string{"create",
						"-d",
						"virtualbox",
						"--virtualbox-memory", "4096",
						"--virtualbox-cpu-count", "2",
						"manager1"}

	fmt.Fprintln(os.Stderr, "Creating manager1 machine ...")
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}
}