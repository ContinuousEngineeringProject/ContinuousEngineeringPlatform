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

func TestDmCreateToReturnRunningNode(t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		nodeStatus := DmCreate(testNodeData[i], testNodeData[i].PREFIX + strconv.Itoa(i+1))
		if nodeStatus != "Running" {
			t.Error("For", testNodeData[i].PREFIX + strconv.Itoa(i+1), "expected Running got", nodeStatus,)
		}
	}
}

func TestDmStatus(t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		nodeStatus := dmStatus(testNodeData[i].PREFIX + strconv.Itoa(i+1))
		if nodeStatus != "Running" {
			t.Error("For", testNodeData[i].PREFIX + strconv.Itoa(i+1), "expected Running got", nodeStatus,)
		}
	}
}

func TestDmStop (t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		nodeStatus := DmStop(testNodeData[i].PREFIX + strconv.Itoa(i+1))
		if nodeStatus != "Stopped" {
			t.Error("For", testNodeData[i].PREFIX + strconv.Itoa(i+1), "expected Stopped got", nodeStatus,)
		}
	}
}

func TestDmStart (t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		nodeStatus := DmStart(testNodeData[i].PREFIX + strconv.Itoa(i+1))
		if nodeStatus != "Running" {
			t.Error("For", testNodeData[i].PREFIX + strconv.Itoa(i+1), "expected RUNNING got", nodeStatus,)
		}
	}
}

func TestDmSSH(t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
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
	for i := 0; i < len(testNodeData); i++ {
		nodeStatus := DmRemove(testNodeData[i].PREFIX + strconv.Itoa(i+1))
		if nodeStatus != "REMOVED" {
			t.Error("For", testNodeData[i].PREFIX + strconv.Itoa(i+1), "expected REMOVED got", nodeStatus,)
		}
	}
}