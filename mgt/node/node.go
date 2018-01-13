package node

import (
	"fmt"
	"github.com/ContinuousEngineeringProject/cePlatform/docker/bash"
	"os"
	"strconv"
)

type ListNodeStatus struct {
	NODE string
	STATUS string
}

// CreateNodes will create multiple nodes
//
func CreateNodes(nodeArgs bash.ArgsCreateNode) (nodesStatus []ListNodeStatus){
	fmt.Fprintln(os.Stderr, "Creating "+strconv.Itoa(nodeArgs.COUNT)+" nodes")

	nodesStatus = make([]ListNodeStatus,nodeArgs.COUNT)
	if nodeArgs.COUNT > 0 {
		//Populate node names
		for n := 0; n < nodeArgs.COUNT; n++ {
			nodesStatus[n].NODE = nodeArgs.PREFIX + strconv.Itoa(n+1)
		}
		//Create the nodes & report status
		for n := 0; n < len(nodesStatus); n++ {
			nodesStatus[n].STATUS = bash.DmCreate(nodeArgs, nodesStatus[n].NODE)
		}
	} else {
		fmt.Fprintln(os.Stderr, "Error Creating node: "+strconv.Itoa(nodeArgs.COUNT)+" is an invalid number of nodes")
		nodesStatus[0].NODE = "NODE_COUNT"
		nodesStatus[0].STATUS = "INVALID"
	}
	return nodesStatus
}

// StopNodes will stop multiple nodes
//
func StopNodes(nodeNames []string) (nodesStatus []ListNodeStatus){
	fmt.Fprintln(os.Stderr, "Stopping "+strconv.Itoa(len(nodeNames))+" nodes")

	for node := 0; node < len(nodeNames); node++ {
		bash.DmStop(nodeNames[node])
	}
	return nodesStatus
}

// StartNodes will start multiple stopped nodes
//
func StartNodes(nodeNames []string) (nodesStatus []ListNodeStatus){
	fmt.Fprintln(os.Stderr, "Starting "+strconv.Itoa(len(nodeNames))+" nodes")

	for node := 0; node < len(nodeNames); node++ {
		bash.DmStart(nodeNames[node])
	}
	return nodesStatus
}

// RemoveNodes will start multiple stopped nodes
//
func RemoveNodes(nodeNames []string) (nodesStatus []ListNodeStatus){
	fmt.Fprintln(os.Stderr, "Removing "+strconv.Itoa(len(nodeNames))+" nodes")

	for node := 0; node < len(nodeNames); node++ {
		bash.DmRemove(nodeNames[node])
	}
	return nodesStatus
}
