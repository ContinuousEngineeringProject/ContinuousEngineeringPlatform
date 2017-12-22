package bash

import (
	"fmt"
	"os"
	"bufio"
	"os/exec"
)

const dCmdName = "docker-machine"

type ArgsVirtualbox struct {
	NAME 	string
	DRIVER 	string
	MEMORY 	string
	CPU 	string
}

func CreateVM(options ArgsVirtualbox) (*exec.Cmd) {
	dCmd := "create"
	switch options.DRIVER {
	case "virtualbox":
		dCmdAgrs := []string{dCmd, "-d", options.DRIVER, "--virtualbox-memory", options.MEMORY, "--virtualbox-cpu-count", options.CPU, options.NAME}
		fmt.Fprintln(os.Stderr, "Creating virtual machine " + options.NAME)
		return exec.Command(dCmdName, dCmdAgrs...)
	default:
		fmt.Fprintln(os.Stderr, "Error unknown driver", options.DRIVER)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "Error creating VM", options.NAME)
	return exec.Command(dCmdName)
}

func RunCmd(cmd *exec.Cmd)  {
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("docker build out | %s\n", scanner.Text())
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