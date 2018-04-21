package node

import (
	"testing"
	"github.com/ContinuousEngineeringProject/cePlatform/docker/bash"
	"strconv"
)

func createTestNodeData() []bash.ArgsCreateNode {
	var argsTestNodes = []bash.ArgsCreateNode{
		{"TestVBNode","virtualbox","1024","1",2},
	}
	return argsTestNodes
}

func createTestNodeNameList(testNodeData bash.ArgsCreateNode) (testNodeNames []string) {
	testNodeNames = make([]string,testNodeData.COUNT)
	for i := 0; i < testNodeData.COUNT; i++ {
		testNodeNames[i] = testNodeData.PREFIX + strconv.Itoa(i+1)
	}
	return testNodeNames
}

func TestCreateNodesToReturnMultipleNodesAreCreatedAndRunning(t *testing.T) {
	testNodeData := createTestNodeData()

	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		nodesStatus := CreateNodes(testNodeData[testIteration])

		//Validate
		for node :=0; node < len(nodesStatus); node++ {
			if nodesStatus[node].STATUS != "Running" {
				t.Error("For", nodesStatus[node].NODE, "expected Running got", nodesStatus[node].STATUS, )
			}
		}
	}
}

func TestStopNodesToReturnMultipleNodesAreStopped(t *testing.T) {
	testNodeData := createTestNodeData()

	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		nodesStatus := StopNodes(createTestNodeNameList(testNodeData[testIteration]))

		//Validate
		for node := 0; node < len(nodesStatus); node++ {
			if nodesStatus[node].STATUS != "Stopped" {
				t.Error("For", nodesStatus[node].NODE, "expected Stopped got", nodesStatus[node].STATUS, )
			}
		}
	}
}

func TestStartNodes(t *testing.T) {
	testNodeData := createTestNodeData()

	for i := 0; i < len(testNodeData); i++ {
		nodeNames := createTestNodeNameList(testNodeData[i])
		nodesStatus := StartNodes(nodeNames)

		//Validate
		for s:=0; s < len(nodesStatus); s++ {
			if nodesStatus[s].STATUS != "RUNNING" {
				t.Error("For", nodesStatus[s].NODE, "expected RUNNING got", nodesStatus[s].STATUS, )
			}
		}
	}
}

func TestRemoveNodesToReturnMultipleNodesAreDeleted(t *testing.T) {
	testNodeData := createTestNodeData()

	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		nodesStatus := RemoveNodes(createTestNodeNameList(testNodeData[testIteration]))

		//Validate
		for node :=0; node < len(nodesStatus); node++ {
			if nodesStatus[node].STATUS != "Successfully removed "+nodesStatus[node].NODE {
				t.Error("For", nodesStatus[node].NODE, "expected", "'Successfully removed "+nodesStatus[node].NODE+"'", "got", "'"+nodesStatus[node].STATUS+"'", )
			}
		}
	}
}
