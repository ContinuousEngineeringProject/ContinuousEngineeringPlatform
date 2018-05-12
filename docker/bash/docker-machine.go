package bash

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

const dmCmd = "docker-machine"

// DmCreate will create a node using docker-machine
//
func DmCreate(options ArgCreateNode, nodeName string) (status string){
	dCmd := "create"
	// Build arguments based on specified driver
	switch options.DRIVER {
	case "virtualbox":
		dCmdAgrs := []string{dCmd, "-d", options.DRIVER, "--virtualbox-memory", options.MEMORY, "--virtualbox-cpu-count", options.CPU, nodeName}
		fmt.Fprintln(os.Stderr, "Creating node "+nodeName+"...")
		runBashCmd(exec.Command(dmCmd, dCmdAgrs...))

		return DmStatus(nodeName)
	default:
		// Throw unknown driver error
		fmt.Fprintln(os.Stderr, "Error unknown driver", options.DRIVER)

		return "Error"
	}
}

// DmStatus will return the status of a node
//
func DmStatus(nodeName string) (nodeStatus string){
	dCmd := "status"
	dCmdArgs := []string{dCmd,nodeName}

	return runBashCmd(exec.Command(dmCmd, dCmdArgs...))
}

// DmRemove will remove a node
//
func DmRemove(nodeName string) (status string){
	dCmd := "rm"
	dCmdAgrs := []string{dCmd, nodeName, "--force"}

	fmt.Fprintln(os.Stderr, "Removing node "+nodeName+"...")

	// TODO: refactor the response to be "Removed"
	return runBashCmd(exec.Command(dmCmd, dCmdAgrs...))
}

// DmStop will stop a running node
//
func DmStop(nodeName string) (status string){
	dCmd := "stop"
	dCmdAgrs := []string{dCmd, nodeName}

	fmt.Fprintln(os.Stderr, "Stoping node "+nodeName+"...")
	runBashCmd(exec.Command(dmCmd, dCmdAgrs...))

	return DmStatus(nodeName)
}

// DmStart will start a stopped node
//
func DmStart(nodeName string) (status string){
	dCmd := "start"
	dCmdAgrs := []string{dCmd, nodeName}

	fmt.Fprintln(os.Stderr, "Starting node "+nodeName+"...")
	runBashCmd(exec.Command(dmCmd, dCmdAgrs...))

	return DmStatus(nodeName)
}

// DmSSH will ssh to a selected node and run the specified command
//
func DmSSH(nodeName string, bashCmd string) (sshOutput string){
	dCmd := "ssh"
	dCmdArgs := []string{dCmd,nodeName,strconv.Quote(bashCmd)}

	fmt.Fprintln(os.Stderr, "SSH to node "+nodeName+"...")

	return runBashCmd(exec.Command(dmCmd, dCmdArgs...))
}

/*
// DmSCP will copy files or directories between nodes and hosts
//
func DmSCP(locationSource string, locationDestination string, isFile bool) (scpStatus string){
	dCmd := "scp"
	if isFile==true {
		dCmdArgs := []string{dCmd, locationSource, locationDestination}
		scpStatus = runBashCmd(exec.Command(dmCmd, dCmdArgs...))
	} else {
		dCmdArgs := []string{dCmd, "-r", locationSource, locationDestination}
		scpStatus = runBashCmd(exec.Command(dmCmd, dCmdArgs...))
	}
	// TODO: Verify that the command ran & Return the the output
	return "EXEC"
}
*/

// DmRestart will restart a node
//
func DmRestart(nodeName string) (status string){
	dCmd := "restart"
	dCmdAgrs := []string{dCmd, nodeName}

	fmt.Fprintln(os.Stderr, "Restarting node "+nodeName+"...")
	runBashCmd(exec.Command(dmCmd, dCmdAgrs...))

	return DmStatus(nodeName)
}

// DmIp will return the IP of a node
//
func DmIp(nodeName string) (nodeIp string){
	dCmd := "ip"
	dCmdAgrs := []string{dCmd, nodeName}

	nodeIp = runBashCmd(exec.Command(dmCmd, dCmdAgrs...))
//	fmt.Fprintln(os.Stderr, nodeName+" IP is "+nodeIp)

	return
}