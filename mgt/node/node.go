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
func CreateNodes(nodeConfig bash.ArgsCreateNode) (nodesStatus []ListNodeStatus){
	fmt.Fprintln(os.Stderr, "Creating "+strconv.Itoa(nodeConfig.COUNT)+" nodes")

	nodesStatus = make([]ListNodeStatus, nodeConfig.COUNT)
	if nodeConfig.COUNT > 0 {
		//Populate node names
		for n := 0; n < nodeConfig.COUNT; n++ {
			nodesStatus[n].NODE = nodeConfig.PREFIX + strconv.Itoa(n+1)
		}
		//Create the nodes & report status
		for n := 0; n < len(nodesStatus); n++ {
			nodesStatus[n].STATUS = bash.DmCreate(nodeConfig, nodesStatus[n].NODE)
		}
	} else {
		fmt.Fprintln(os.Stderr, "Error Creating node: "+strconv.Itoa(nodeConfig.COUNT)+" is an invalid number of nodes")
		nodesStatus[0].NODE = "NODE_COUNT"
		nodesStatus[0].STATUS = "INVALID"
	}
	return
}

// StopNodes will stop multiple nodes
//
func StopNodes(nodeNames []string) (nodesStatus []ListNodeStatus){
	fmt.Fprintln(os.Stderr, "Stopping "+strconv.Itoa(len(nodeNames))+" nodes")

	nodesStatus = make([]ListNodeStatus, len(nodeNames))
	for node := 0; node < len(nodeNames); node++ {
		nodesStatus[node].NODE = nodeNames[node]
		nodesStatus[node].STATUS = bash.DmStop(nodeNames[node])
	}
	return
}

// StartNodes will start multiple stopped nodes
//
func StartNodes(nodeNames []string) (nodesStatus []ListNodeStatus){
	fmt.Fprintln(os.Stderr, "Starting "+strconv.Itoa(len(nodeNames))+" nodes")

	nodesStatus = make([]ListNodeStatus, len(nodeNames))
	for node := 0; node < len(nodeNames); node++ {
		nodesStatus[node].NODE = nodeNames[node]
		nodesStatus[node].STATUS = bash.DmStart(nodeNames[node])
	}
	return 
}

// RemoveNodes will remove multiple nodes
//
func RemoveNodes(nodeNames []string) (nodesStatus []ListNodeStatus){
	fmt.Fprintln(os.Stderr, "Removing "+strconv.Itoa(len(nodeNames))+" nodes")

	nodesStatus = make([]ListNodeStatus, len(nodeNames))
	for node := 0; node < len(nodeNames); node++ {
		nodesStatus[node].NODE = nodeNames[node]
		nodesStatus[node].STATUS = bash.DmRemove(nodeNames[node])
	}
	return
}
