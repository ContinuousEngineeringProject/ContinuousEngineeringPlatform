package node

import (
	"testing"
	"github.com/ContinuousEngineeringProject/cePlatform/docker/bash"
	"strconv"
)

func createTestNodeData() []bash.ArgCreateNode {
	var argsTestNodes = []bash.ArgCreateNode{
		{"TestVBNode","virtualbox","1024","1",2},
	}
	return argsTestNodes
}

func createTestNodeNameList(testNodeData bash.ArgCreateNode) (testNodeNames []string) {
	testNodeNames = make([]string,testNodeData.COUNT)
	for i := 0; i < testNodeData.COUNT; i++ {
		testNodeNames[i] = testNodeData.PREFIX + strconv.Itoa(i+1)
	}
	return testNodeNames
}

func TestCreateNodesToReturnMultipleNodesAreCreatedAndRunning(t *testing.T) {
	testNodeData := createTestNodeData()

	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		nodeStatusList := CreateNodes(testNodeData[testIteration])

		//Validate
		for node :=0; node < len(nodeStatusList); node++ {
			if nodeStatusList[node].STATUS != "Running" {
				t.Error("For", nodeStatusList[node].NODE, "expected", "Running", "got", nodeStatusList[node].STATUS, )
			}
		}
	}
}

func TestStopNodesToReturnMultipleNodesAreStopped(t *testing.T) {
	testNodeData := createTestNodeData()

	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		nodeStatusList := StopNodes(createTestNodeNameList(testNodeData[testIteration]))

		//Validate
		for node := 0; node < len(nodeStatusList); node++ {
			if nodeStatusList[node].STATUS != "Stopped" {
				t.Error("For", nodeStatusList[node].NODE, "expected", "Stopped", "got", nodeStatusList[node].STATUS, )
			}
		}
	}
}

func TestStartNodesMultipleNodesAreStarted(t *testing.T) {
	testNodeData := createTestNodeData()

	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		nodeStatusList := StartNodes(createTestNodeNameList(testNodeData[testIteration]))

		//Validate
		for node :=0; node < len(nodeStatusList); node++ {
			if nodeStatusList[node].STATUS != "Running" {
				t.Error("For", nodeStatusList[node].NODE, "expected", "Running", "got", nodeStatusList[node].STATUS, )
			}
		}
	}
}

func TestRestartNodesMultipleNodesAreRestarted(t *testing.T) {
	testNodeData := createTestNodeData()

	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		nodeStatusList := RestartNodes(createTestNodeNameList(testNodeData[testIteration]))

		//Validate
		for node :=0; node < len(nodeStatusList); node++ {
			if nodeStatusList[node].STATUS != "Running" {
				t.Error("For", nodeStatusList[node].NODE, "expected", "Running", "got", nodeStatusList[node].STATUS, )
			}
		}
	}
}

func TestRemoveNodesToReturnMultipleNodesAreDeleted(t *testing.T) {
	testNodeData := createTestNodeData()

	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		nodeStatusList := RemoveNodes(createTestNodeNameList(testNodeData[testIteration]))

		//Validate
		for node :=0; node < len(nodeStatusList); node++ {
			if nodeStatusList[node].STATUS != "Successfully removed "+nodeStatusList[node].NODE {
				t.Error("For", nodeStatusList[node].NODE, "expected", "'Successfully removed "+nodeStatusList[node].NODE+"'", "got", "'"+nodeStatusList[node].STATUS+"'", )
			}
		}
	}
}
