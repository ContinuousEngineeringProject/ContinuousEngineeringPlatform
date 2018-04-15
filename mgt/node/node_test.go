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

func TestCreateNodesToReturnMultipleRunningNodes(t *testing.T) {
	testNodeData := createTestNodeData()

	for node := 0; node < len(testNodeData); node++ {
		status := CreateNodes(testNodeData[node])

		//Validate
		for i:=0; i < len(status); i++ {
			if status[node].STATUS != "Running" {
				t.Error("For", status[node].NODE, "expected Running got", status[node].STATUS, )
			}
		}
	}
}

func TestStopNodes(t *testing.T) {
	testNodeData := createTestNodeData()

	for i := 0; i < len(testNodeData); i++ {
		nodeNames := createTestNodeNameList(testNodeData[i])
		status := StopNodes(nodeNames)

		//Validate
		for s:=0; s < len(status); s++ {
			if status[s].STATUS != "STOPPED" {
				t.Error("For", status[s].NODE, "expected STOPPED got", status[s].STATUS, )
			}
		}
	}
}

func TestStartNodes(t *testing.T) {
	testNodeData := createTestNodeData()

	for i := 0; i < len(testNodeData); i++ {
		nodeNames := createTestNodeNameList(testNodeData[i])
		status := StartNodes(nodeNames)

		//Validate
		for s:=0; s < len(status); s++ {
			if status[s].STATUS != "RUNNING" {
				t.Error("For", status[s].NODE, "expected RUNNING got", status[s].STATUS, )
			}
		}
	}
}

func TestRemoveNodes(t *testing.T) {
	testNodeData := createTestNodeData()

	for i := 0; i < len(testNodeData); i++ {
		nodeNames := createTestNodeNameList(testNodeData[i])
		status := RemoveNodes(nodeNames)

		//Validate
		for s:=0; s < len(status); s++ {
			if status[s].STATUS != "REMOVED" {
				t.Error("For", status[s].NODE, "expected REMOVED got", status[s].STATUS, )
			}
		}
	}
}
