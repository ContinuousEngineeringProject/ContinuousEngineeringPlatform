package bash

import (
	"fmt"
	"os"
	"os/exec"
)

const cmdName = "docker-machine"

type ArgsCreateNode struct {
	PREFIX string
	DRIVER string
	MEMORY string
	CPU    string
	COUNT  int
}

// DmCreate will create a node using docker-machine
//
func DmCreate(options ArgsCreateNode, nodeName string) (status string){
	dCmd := "create"
	// Build arguments based on specified driver
	switch options.DRIVER {
	case "virtualbox":
		dCmdAgrs := []string{dCmd, "-d", options.DRIVER, "--virtualbox-memory", options.MEMORY, "--virtualbox-cpu-count", options.CPU, nodeName}
		fmt.Fprintln(os.Stderr, "Creating node "+nodeName+"...")
		runBashCmd(exec.Command(cmdName, dCmdAgrs...))
		// TODO: Refactor DmCreate to use actual docker-machine status
		status = "RUNNING"
		fmt.Fprintln(os.Stderr, "Node "+nodeName+" created")
	default:
		// Throw unknown driver error
		status = "ERROR"
		fmt.Fprintln(os.Stderr, "Error unknown driver", options.DRIVER)
	}
	return status
}

// DmRemove will remove a node
//
func DmRemove(nodeName string) (status string){
	dCmd := "rm"
	dCmdAgrs := []string{dCmd, nodeName, "--force"}

	fmt.Fprintln(os.Stderr, "Removing node "+nodeName+"...")
	runBashCmd(exec.Command(cmdName, dCmdAgrs...))
	fmt.Fprintln(os.Stderr, "Node "+nodeName+" removed")
	// TODO: Refactor DmRemove to use docker-machine to confirm node does not exist
	return "REMOVED"
}

// DmStop will stop a running node
//
func DmStop(nodeName string) (status string){
	dCmd := "stop"
	dCmdAgrs := []string{dCmd, nodeName}

	fmt.Fprintln(os.Stderr, "Stoping node "+nodeName+"...")
	runBashCmd(exec.Command(cmdName, dCmdAgrs...))
	fmt.Fprintln(os.Stderr, "Node "+nodeName+" stopped")
	// TODO: Refactor DmStop to use actual docker-machine status
	return "STOPPED"
}

// DmStart will start a stopped node
//
func DmStart(nodeName string) (status string){
	dCmd := "start"
	dCmdAgrs := []string{dCmd, nodeName}

	fmt.Fprintln(os.Stderr, "Starting node "+nodeName+"...")
	runBashCmd(exec.Command(cmdName, dCmdAgrs...))
	fmt.Fprintln(os.Stderr, "Node "+nodeName+" started")
	// TODO: Refactor DmStart to use actual docker-machine status
	return "RUNNING"
}

