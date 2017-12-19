package main

import (
	"fmt"
	"os"
	"bufio"
	"os/exec"
//	"strconv"
//	"github.com/ContinuousEngineeringProject/cePlatform/docker/bash"
)

func createVM(DRIVER string, MEMORY string, CPU string, NAME string) *exec.Cmd {
	switch DRIVER {
	case "virtualbox":
		dmCmd := []string{"create", "-d", DRIVER, "--virtualbox-memory", MEMORY, "--virtualbox-cpu-count", CPU, NAME}
		fmt.Fprintln(os.Stderr, "Creating manager1 machine ...")
		return exec.Command("docker-machine", dmCmd...)
	default:
		fmt.Fprintln(os.Stderr, "Error unknown driver", DRIVER)
		os.Exit(1)
	}
	dmCmd := ""
	return exec.Command(dmCmd)
}

func main() {

	cmd := createVM("virtualbox","4096","2","manager1")


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