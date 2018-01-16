package bash

import "testing"

func createTestNodeData() []ArgsCreateNode {
	var argsTestNodes = []ArgsCreateNode{
		{"TestVBNode","virtualbox","1024","1",1},
//		{"TestVBNode2","virtualbox","1024","1",1},
	}
	return argsTestNodes
}

func TestDmCreate(t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		status := DmCreate(testNodeData[i], testNodeData[i].PREFIX)
		if status != "RUNNING" {
			t.Error("For", testNodeData[i].PREFIX, "expected RUNNING got", status,)
		}
	}
}

func TestDmStop (t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		status := DmStop(testNodeData[i].PREFIX)
		if status != "STOPPED" {
			t.Error("For", testNodeData[i].PREFIX, "expected STOPPED got", status,)
		}
	}
}

func TestDmStart (t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		status := DmStart(testNodeData[i].PREFIX)
		if status != "RUNNING" {
			t.Error("For", testNodeData[i].PREFIX, "expected RUNNING got", status,)
		}
	}
}

func TestDmRemove (t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		status := DmRemove(testNodeData[i].PREFIX)
		if status != "REMOVED" {
			t.Error("For", testNodeData[i].PREFIX, "expected REMOVED got", status,)
		}
	}
}

func TestDmSSH(t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		node := DmCreate(testNodeData[i], testNodeData[i].PREFIX)
		if node == "RUNNING" {
			sshOutput := DmSSH(testNodeData[i].PREFIX,"ls")
			if sshOutput != "EXEC" { //needs to be the expected return value from the ssh
				t.Error("Failed to ssh to ", testNodeData[i].PREFIX, "expected EXEC got", sshOutput,)
			}
		} else {
			t.Error("Failed to create", testNodeData[i].PREFIX,)
		}
	}
	// Remove node(s) created during the test
	for i := 0; i < len(testNodeData); i++ {
		DmRemove(testNodeData[i].PREFIX)
	}
}
