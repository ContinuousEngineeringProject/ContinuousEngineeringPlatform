package bash

import (
	"os/exec"
)

const dsCmd = "docker swarm"


// DsInit will initiate a swarm
//
func DsInit (nodeDetails NodeDetails) (swarmStatus bool){
	dCmd := "init"
	dCmdAgrs := []string{dCmd}

	runBashCmd(exec.Command(dsCmd, dCmdAgrs...))

	return
}