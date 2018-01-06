package main

import (
	"fmt"
	"github.com/ContinuousEngineeringProject/cePlatform/docker/bash"
	"github.com/ContinuousEngineeringProject/cePlatform/mgt/node"
	"os"
)

func main() {

	fmt.Fprintln(os.Stderr, "Starting Legion")

	dmCreateArgs := bash.ArgsCreateNode{"TestNode", "virtualbox", "1024", "1", 2}

	// Create nodes
	node.CreateNodes(dmCreateArgs)

	// Stop nodes
	nodeNames := []string{"TestNode1","TestNode2"}
	node.StopNodes(nodeNames)

	// Start nodes
	node.StartNodes(nodeNames)

	// Remove nodes
	node.RemoveNodes(nodeNames)

}
