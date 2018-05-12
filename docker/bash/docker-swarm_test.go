package bash

import (
	"testing"
)

func createTestSwarmData() []NodeDetails {
	var swarmTestNodes = []NodeDetails{
		{"TestVBNode1", DmIp("TestVBNode1") , DmStatus("TestVBNode1"), "Master", false,},
//		{"TestVBNode2", DmIp("TestVBNode2") , DmStatus("TestVBNode2"), "Worker", false,},
//		{"TestVBNode3", DmIp("TestVBNode3") , DmStatus("TestVBNode3"), "Worker", false,},
	}

	return swarmTestNodes
}

func TestDsInitToReturnInitialisedSwarmWithSingleManager(t *testing.T) {
	swarmTestData := createTestSwarmData()
	for node := 0; node < len(swarmTestData); node++ {
		if swarmTestData[node].ROLE == "Master" && swarmTestData[node].STATUS == "Running"{
			swarmTestData[node].SWARM = DsInit(swarmTestData[node])
			if swarmTestData[node].SWARM != true {
				t.Error("For", swarmTestData[node].NODE, "expected", "true", "got", swarmTestData[node].SWARM, )
			}
		}
	}
}