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
	dCmd := "create"
	// Build arguments based on specified driver
	switch options.DRIVER {
	case "virtualbox":
		dCmdAgrs := []string{dCmd, "-d", options.DRIVER, "--virtualbox-memory", options.MEMORY, "--virtualbox-cpu-count", options.CPU, nodeName}
		fmt.Fprintln(os.Stderr, "Creating node "+nodeName+"...")
		runBashCmd(exec.Command(cmdName, dCmdAgrs...))

		return dmStatus(nodeName)
	default:
		// Throw unknown driver error
		fmt.Fprintln(os.Stderr, "Error unknown driver", options.DRIVER)

		return "Error"
	}
}

// dmStatus will return the status of a node
//
func dmStatus(nodeName string) (nodeStatus string){
	dCmd := "status"
	dCmdArgs := []string{dCmd,nodeName}

	return runBashCmd(exec.Command(cmdName, dCmdArgs...))
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

// DmSSH will ssh to a selected node and run the specified command
//
func DmSSH(nodeName string, bashCmd string) (sshOutput string){
	dCmd := "ssh"
	dCmdArgs := []string{dCmd,nodeName,strconv.Quote(bashCmd)}
	// TODO: Retrieve the output from the command run on the node
	runBashCmd(exec.Command(cmdName, dCmdArgs...))
	return "EXEC"
}

// DmSCP will copy files or directories between nodes and hosts
//
func DmSCP(locationSource string, locationDestination string, isFile bool) (scpStatus string){
	dCmd := "scp"
	if isFile==true {
		dCmdArgs := []string{dCmd, locationSource, locationDestination}
		runBashCmd(exec.Command(cmdName, dCmdArgs...))
	} else {
		dCmdArgs := []string{dCmd, "-r", locationSource, locationDestination}
		runBashCmd(exec.Command(cmdName, dCmdArgs...))
	}
	// TODO: Verify that the command ran & Return the the output
	return "EXEC"
}