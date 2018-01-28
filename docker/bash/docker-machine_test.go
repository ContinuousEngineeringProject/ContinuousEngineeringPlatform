package bash

import "testing"


func createTestNodeData() []ArgsCreateNode {
	var argsTestNodes = []ArgsCreateNode{
		{"TestVBNode","virtualbox","1024","1",1},
//		{"TestVBNode2","virtualbox","1024","1",1},
	}
	return argsTestNodes
}

// TestDmCreate will test the DmCreate function
//
func TestDmCreate(t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		status := DmCreate(testNodeData[i], testNodeData[i].PREFIX)
		if status != "RUNNING" {
			t.Error("For", testNodeData[i].PREFIX, "expected RUNNING got", status,)
		}
	}
}

// TestDmStop will test the DmStop function
//
func TestDmStop (t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		status := DmStop(testNodeData[i].PREFIX)
		if status != "STOPPED" {
			t.Error("For", testNodeData[i].PREFIX, "expected STOPPED got", status,)
		}
	}
}

// TestDmStart will test the DmStart function
//
func TestDmStart (t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		status := DmStart(testNodeData[i].PREFIX)
		if status != "RUNNING" {
			t.Error("For", testNodeData[i].PREFIX, "expected RUNNING got", status,)
		}
	}
}

// TestDmRemove will test the DmRemove function
//
func TestDmRemove (t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		status := DmRemove(testNodeData[i].PREFIX)
		if status != "REMOVED" {
			t.Error("For", testNodeData[i].PREFIX, "expected REMOVED got", status,)
		}
	}
}

// TestDmSSH will test the DmSSH function
//
func TestDmSSH(t *testing.T) {
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		node := DmCreate(testNodeData[i], testNodeData[i].PREFIX)
		if node == "RUNNING" {
			sshOutput := DmSSH(testNodeData[i].PREFIX,"ls")
			if sshOutput != "EXEC" { //TODO: needs to be the expected return value from the ssh
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

// TestDmSCP will test the DmSCP function
//
func TestDmSCP(t *testing.T) {
	// TODO: Refactor to include multiple source & dest locations
	testNodeData := createTestNodeData()
	for i := 0; i < len(testNodeData); i++ {
		node := DmCreate(testNodeData[i], testNodeData[i].PREFIX)
		if node == "RUNNING" {
			scpStatus := DmSCP("./docker-machine_test.go",testNodeData[i].PREFIX+":~",true)
			if scpStatus != "EXEC" { //TODO: needs to be the expected return value from the ssh
				t.Error("Failed to scp to", testNodeData[i].PREFIX, "expected EXEC got", scpStatus,)
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