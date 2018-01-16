package bash

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
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
	// TODO: Refactor DmCreate to use actual docker-machine status
	dCmd := "create"
	// Build arguments based on specified driver
	switch options.DRIVER {
	case "virtualbox":
		dCmdAgrs := []string{dCmd, "-d", options.DRIVER, "--virtualbox-memory", options.MEMORY, "--virtualbox-cpu-count", options.CPU, nodeName}
		fmt.Fprintln(os.Stderr, "Creating node "+nodeName+"...")
		runBashCmd(exec.Command(cmdName, dCmdAgrs...))
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
	// TODO: Refactor DmRemove to use docker-machine to confirm node does not exist
	dCmd := "rm"
	dCmdAgrs := []string{dCmd, nodeName, "--force"}

	fmt.Fprintln(os.Stderr, "Removing node "+nodeName+"...")
	runBashCmd(exec.Command(cmdName, dCmdAgrs...))
	fmt.Fprintln(os.Stderr, "Node "+nodeName+" removed")
	return "REMOVED"
}

// DmStop will stop a running node
//
func DmStop(nodeName string) (status string){
	// TODO: Refactor DmStop to use actual docker-machine status
	dCmd := "stop"
	dCmdAgrs := []string{dCmd, nodeName}

	fmt.Fprintln(os.Stderr, "Stoping node "+nodeName+"...")
	runBashCmd(exec.Command(cmdName, dCmdAgrs...))
	fmt.Fprintln(os.Stderr, "Node "+nodeName+" stopped")
	return "STOPPED"
}

// DmStart will start a stopped node
//
func DmStart(nodeName string) (status string){
	// TODO: Refactor DmStart to use actual docker-machine status
	dCmd := "start"
	dCmdAgrs := []string{dCmd, nodeName}

	fmt.Fprintln(os.Stderr, "Starting node "+nodeName+"...")
	runBashCmd(exec.Command(cmdName, dCmdAgrs...))
	fmt.Fprintln(os.Stderr, "Node "+nodeName+" started")
	return "RUNNING"
}

// DmSSH will ssh to the node
// and run your specified command on the node
func DmSSH(nodeName string, bashCmd string) (sshOutput string){
	// TODO: Refactor DmSSH
	dCmd := "ssh"
	dCmdArgs := []string{dCmd,nodeName,strconv.Quote(bashCmd)}

	runBashCmd(exec.Command(cmdName, dCmdArgs...))
	return "EXEC"
}

