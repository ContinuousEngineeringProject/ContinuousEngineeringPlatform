package bash

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

const cmdName = "docker-machine"

type ArgsCreateVM struct {
	NAME 	string
	DRIVER 	string
	MEMORY 	string
	CPU 	string
	COUNT 	int
}

func CreateVM(options ArgsCreateVM) {
	dCmd := "create"

	for vm := 1; vm <= options.COUNT; vm++ {
		fmt.Fprintln(os.Stderr, "Creating virtual machine " + options.NAME + strconv.Itoa(vm) + "...")

		switch options.DRIVER {
		case "virtualbox":
			dCmdAgrs := []string{dCmd, "-d", options.DRIVER, "--virtualbox-memory", options.MEMORY, "--virtualbox-cpu-count", options.CPU, options.NAME + strconv.Itoa(vm)}
			runBashCmd(exec.Command(cmdName, dCmdAgrs...))
			fmt.Fprintln(os.Stderr, "Virtual machine " + options.NAME + strconv.Itoa(vm) + " created")
		default:
			fmt.Fprintln(os.Stderr, "Error unknown driver", options.DRIVER)
			os.Exit(1)
		}
	}
	return
}

