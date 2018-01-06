package node

import (
	"fmt"
	"github.com/ContinuousEngineeringProject/cePlatform/docker/bash"
	"os"
	"strconv"
)

// CreateNodes will create multiple nodes
//
func CreateNodes(nodeArgs bash.ArgsCreateNode) {
	fmt.Fprintln(os.Stderr, "Creating "+strconv.Itoa(nodeArgs.COUNT)+" nodes")

	if nodeArgs.COUNT > 0 {
		for node := 1; node <= nodeArgs.COUNT; node++ {
			nodeName := nodeArgs.PREFIX+strconv.Itoa(node)
			bash.DmCreate(nodeArgs, nodeName)
		}
	} else {
		fmt.Fprintln(os.Stderr, "Error Creating node: "+strconv.Itoa(nodeArgs.COUNT)+" is an invalid number of nodes")
	}
	//TODO: Return the created nodes statuses
	return
}

// StopNodes will stop multiple nodes
//
func StopNodes(nodeNames []string) {
	fmt.Fprintln(os.Stderr, "Stopping "+strconv.Itoa(len(nodeNames))+" nodes")

	for node := 0; node <= len(nodeNames)-1; node++ {
		bash.DmStop(nodeNames[node])
	}
	//TODO: Return the nodes statuses
	return
}

// StartNodes will start multiple stopped nodes
//
func StartNodes(nodeNames []string){
	fmt.Fprintln(os.Stderr, "Starting "+strconv.Itoa(len(nodeNames))+" nodes")

	for node := 0; node <= len(nodeNames)-1; node++ {
		bash.DmStart(nodeNames[node])
	}
	//TODO: Return the nodes statuses
	return
}

// RemoveNodes will start multiple stopped nodes
//
func RemoveNodes(nodeNames []string){
	fmt.Fprintln(os.Stderr, "Removing "+strconv.Itoa(len(nodeNames))+" nodes")

	for node := 0; node <= len(nodeNames)-1; node++ {
		bash.DmRemove(nodeNames[node])
	}
	//TODO: Return the nodes statuses
	return
}


