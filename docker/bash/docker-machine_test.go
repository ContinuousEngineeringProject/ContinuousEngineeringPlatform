package bash

import "testing"

func createTestNodeData() []ArgsCreateNode {
	var argsTestNodes = []ArgsCreateNode{
		{"TestVBNode1","virtualbox","1024","1",1},
//		{"TestVBNode2","virtualbox","1024","1",1},
	}
	return argsTestNodes
}

func TestDmCreate(t *testing.T) {
	testNodeData := createTestNodeData()
	for node := 0; node < len(testNodeData); node++ {
		status := DmCreate(testNodeData[node], testNodeData[node].PREFIX)
		if status != "RUNNING" {
			t.Error("For", testNodeData[node].PREFIX, "expected RUNNING got", status,)
		}
	}
}

func TestDmStop (t *testing.T) {
	testNodeData := createTestNodeData()
	for node := 0; node < len(testNodeData); node++ {
		status := DmStop(testNodeData[node].PREFIX)
		if status != "STOPPED" {
			t.Error("For", testNodeData[node].PREFIX, "expected STOPPED got", status,)
		}
	}
}

func TestDmStart (t *testing.T) {
	testNodeData := createTestNodeData()
	for node := 0; node < len(testNodeData); node++ {
		status := DmStart(testNodeData[node].PREFIX)
		if status != "RUNNING" {
			t.Error("For", testNodeData[node].PREFIX, "expected RUNNING got", status,)
		}
	}
}

func TestDmRemove (t *testing.T) {
	testNodeData := createTestNodeData()
	for node := 0; node < len(testNodeData); node++ {
		status := DmRemove(testNodeData[node].PREFIX)
		if status != "REMOVED" {
			t.Error("For", testNodeData[node].PREFIX, "expected REMOVED got", status,)
		}
	}
}