package bash

import (
	"testing"
	"strconv"
)

func removeTestNodes(testNodeData []ArgsCreateNode){
	for i := 0; i < len(testNodeData); i++ {
		DmRemove(testNodeData[i].PREFIX)
	}
}

func createTestNodeData() []ArgsCreateNode {
	var argsTestNodes = []ArgsCreateNode{
		{"TestVBNode","virtualbox","1024","1",1},
	}
	return argsTestNodes
}

func TestDmCreateToReturnCreateSingleRunningNode(t *testing.T) {
	testNodeData := createTestNodeData()
	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		for node := 0; node < testNodeData[testIteration].COUNT; node++ {
			nodeStatus := DmCreate(testNodeData[testIteration], testNodeData[testIteration].PREFIX + strconv.Itoa(node+1))
			if nodeStatus != "Running" {
				t.Error("For", testNodeData[testIteration].PREFIX + strconv.Itoa(node+1), "expected Running got", nodeStatus,)
			}
		}
	}
}

func TestDmStatusToReturnStatusOfSingleRunningNode(t *testing.T) {
	testNodeData := createTestNodeData()
	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		for node := 0; node < testNodeData[testIteration].COUNT; node++ {
			nodeStatus := DmStatus(testNodeData[testIteration].PREFIX + strconv.Itoa(node+1))
			if nodeStatus != "Running" {
				t.Error("For", testNodeData[testIteration].PREFIX + strconv.Itoa(node+1), "expected Running got", nodeStatus,)
			}
		}
	}
}

func TestDmStopToReturnSingleNodeIsStopped (t *testing.T) {
	testNodeData := createTestNodeData()
	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		for node := 0; node < testNodeData[testIteration].COUNT; node++ {
			nodeStatus := DmStop(testNodeData[node].PREFIX + strconv.Itoa(node+1))
			if nodeStatus != "Stopped" {
				t.Error("For", testNodeData[testIteration].PREFIX + strconv.Itoa(node+1), "expected Stopped got", nodeStatus,)
			}
		}
	}
}

func TestDmStart (t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		for n := 0; n < testNodeData[i].COUNT; n++ {

		}

		nodeStatus := DmStart(testNodeData[i].PREFIX + strconv.Itoa(i+1))
		if nodeStatus != "Running" {
			t.Error("For", testNodeData[i].PREFIX + strconv.Itoa(i+1), "expected RUNNING got", nodeStatus,)
		}
	}
}

func TestDmSSH(t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		for n := 0; n < testNodeData[i].COUNT; n++ {

		}

		nodeStatus := DmCreate(testNodeData[i], testNodeData[i].PREFIX + strconv.Itoa(i+1))
		if nodeStatus == "Running" {
			sshOutput := DmSSH(testNodeData[i].PREFIX + strconv.Itoa(i+1),"ls")
			if sshOutput != "EXEC" { //TODO: needs to be the expected return value from the ssh
				t.Error("Failed to ssh to ", testNodeData[i].PREFIX + strconv.Itoa(i+1), "expected EXEC got", sshOutput,)
			}
		} else {
			t.Error("Failed to create", testNodeData[i].PREFIX + strconv.Itoa(i+1),)
		}
	}
	// Remove node(s) created during the test
	removeTestNodes(testNodeData)
}

func TestDmSCP(t *testing.T) {
	// TODO: Refactor to include multiple source & dest locations
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		for n := 0; n < testNodeData[i].COUNT; n++ {

		}

		nodeStatus := DmCreate(testNodeData[i], testNodeData[i].PREFIX + strconv.Itoa(i+1))
		if nodeStatus == "Running" {
			scpStatus := DmSCP("./docker-machine_test.go",testNodeData[i].PREFIX+strconv.Itoa(i+1)+":~",true)
			if scpStatus != "EXEC" { //TODO: needs to be the expected return value from the ssh
				t.Error("Failed to scp to", testNodeData[i].PREFIX + strconv.Itoa(i+1), "expected EXEC got", scpStatus,)
			}
		} else {
			t.Error("Failed to create", testNodeData[i].PREFIX + strconv.Itoa(i+1),)
		}
	}
	// Remove node(s) created during the test
	removeTestNodes(testNodeData)
}

func TestDmRemove (t *testing.T) {
	testNodeData := createTestNodeData()
	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		for node := 0; node < testNodeData[testIteration].COUNT; node++ {
			nodeName := testNodeData[testIteration].PREFIX + strconv.Itoa(node+1)
			nodeStatus := DmRemove(nodeName)
			if nodeStatus != "Successfully removed "+nodeName {
				t.Error("For", nodeName, "expected 'Successfully removed "+nodeName+"' got", "'"+nodeStatus+"'",)
			}
		}
	}
}