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
func DmCreate(options ArgsCreateNode, nodeName string) {
	dCmd := "create"

	// Build arguments based on specified driver
	switch options.DRIVER {
	case "virtualbox":
		dCmdAgrs := []string{dCmd, "-d", options.DRIVER, "--virtualbox-memory", options.MEMORY, "--virtualbox-cpu-count", options.CPU, nodeName}
		fmt.Fprintln(os.Stderr, "Creating node "+nodeName+"...")
		runBashCmd(exec.Command(cmdName, dCmdAgrs...))
		fmt.Fprintln(os.Stderr, "Node "+nodeName+" created")
	default:
		// Throw unknown driver error
		fmt.Fprintln(os.Stderr, "Error unknown driver", options.DRIVER)
	}
	//TODO: Return the created node status
	return
}

// DmRemove will remove a node
//
func DmRemove(nodeName string) {
	dCmd := "rm"
	dCmdAgrs := []string{dCmd, nodeName, "--force"}

	fmt.Fprintln(os.Stderr, "Removing node "+nodeName+"...")
	runBashCmd(exec.Command(cmdName, dCmdAgrs...))
	fmt.Fprintln(os.Stderr, "Node "+nodeName+" removed")
	//TODO: Return the node status
	return
}

// DmStop will stop a running node
//
func DmStop(nodeName string) {
	dCmd := "stop"
	dCmdAgrs := []string{dCmd, nodeName}

	fmt.Fprintln(os.Stderr, "Stoping node "+nodeName+"...")
	runBashCmd(exec.Command(cmdName, dCmdAgrs...))
	fmt.Fprintln(os.Stderr, "Node "+nodeName+" stopped")
	//TODO: Return the node status
	return
}

// DmStart will start a stopped node
//
func DmStart(nodeName string) {
	dCmd := "start"
	dCmdAgrs := []string{dCmd, nodeName}

	fmt.Fprintln(os.Stderr, "Starting node "+nodeName+"...")
	runBashCmd(exec.Command(cmdName, dCmdAgrs...))
	fmt.Fprintln(os.Stderr, "Node "+nodeName+" started")
	//TODO: Return the node status
	return
}
