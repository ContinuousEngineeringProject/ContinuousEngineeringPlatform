package bash

import (
	"testing"
	"strconv"
)

func removeTestNodes(testNodeData []ArgCreateNode){
	for i := 0; i < len(testNodeData); i++ {
		DmRemove(testNodeData[i].PREFIX)
	}
}

func createTestNodeData() []ArgCreateNode {
	var argsTestNodes = []ArgCreateNode{
		{"TestVBNode","virtualbox","1024","1",1},
	}
	return argsTestNodes
}

func TestDmCreateToReturnCreateRunningNode(t *testing.T) {
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

func TestDmStatusToReturnStatusOfRunningNode(t *testing.T) {
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

func TestDmStopToReturnNodeIsStopped (t *testing.T) {
	testNodeData := createTestNodeData()
	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		for node := 0; node < testNodeData[testIteration].COUNT; node++ {
			nodeStatus := DmStop(testNodeData[testIteration].PREFIX + strconv.Itoa(node+1))
			if nodeStatus != "Stopped" {
				t.Error("For", testNodeData[testIteration].PREFIX + strconv.Itoa(node+1), "expected Stopped got", nodeStatus,)
			}
		}
	}
}

func TestDmStartToReturnNodeIsStarted (t *testing.T) {
	testNodeData := createTestNodeData()
	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		for node := 0; node < testNodeData[testIteration].COUNT; node++ {
			nodeStatus := DmStart(testNodeData[testIteration].PREFIX + strconv.Itoa(node+1))
			if nodeStatus != "Running" {
				t.Error("For", testNodeData[testIteration].PREFIX + strconv.Itoa(node+1), "expected Running got", nodeStatus,)
			}
		}
	}
}

func TestDmRestartToReturnNodeIsRestarted(t *testing.T) {
	testNodeData := createTestNodeData()
	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		for node := 0; node < testNodeData[testIteration].COUNT; node++ {
			nodeStatus := DmRestart(testNodeData[testIteration].PREFIX + strconv.Itoa(node+1))
			if nodeStatus != "Running" {
				t.Error("For", testNodeData[testIteration].PREFIX + strconv.Itoa(node+1), "expected Running got", nodeStatus,)
			}
		}
	}
}

func TestDmSSHToReturnSshToNode(t *testing.T) {
	testNodeData := createTestNodeData()
	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		for node := 0; node < testNodeData[testIteration].COUNT; node++ {
			sshOutput := DmSSH(testNodeData[testIteration].PREFIX + strconv.Itoa(node+1),"echo test")
			if sshOutput != "test" {
				t.Error("Failed to ssh to", testNodeData[testIteration].PREFIX + strconv.Itoa(node+1), "expected", "test", "got", sshOutput,)
			}
		}
	}
}

/*
func TestDmSCP(t *testing.T) {
	// TODO: Refactor to include multiple source & dest locations
	testNodeData := createTestSwarmData()
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
*/

func TestDmIpToReturnNodeIp(t *testing.T) {
	testNodeData := createTestNodeData()
	for testIteration := 0; testIteration < len(testNodeData); testIteration++ {
		for node := 0; node < testNodeData[testIteration].COUNT; node++ {
			nodeIp := DmIp(testNodeData[testIteration].PREFIX + strconv.Itoa(node+1))
			if nodeIp != "192.168.99.100" {
				t.Error("For", testNodeData[testIteration].PREFIX + strconv.Itoa(node+1), "expected", "192.168.99.100", "got", nodeIp,)
			}
		}
	}
}

func TestDmRemoveToReturnNodeIsRemoved (t *testing.T) {
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