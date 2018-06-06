package bash

import (
	"fmt"
	"os"
)

const dsCmd = "docker swarm"


// DsInit will initiate a swarm
//
func DsInit (nodeDetails NodeDetails) (swarmStatus bool){
	dCmd := "init"
//	dCmdAgrs := []string{dCmd, "--listen-addr", nodeDetails.IP, "--advertise-addr", nodeDetails.IP,}
//	runBashCmd(exec.Command(dsCmd, dCmdAgrs...))

	dCmdString := dsCmd+ " " +dCmd+ " --listen-addr " +nodeDetails.IP+ " --advertise-addr " +nodeDetails.IP


	swarmOutput := DmSSH(nodeDetails.NODE, dCmdString)

	fmt.Fprintln(os.Stderr,swarmOutput)

	return
}