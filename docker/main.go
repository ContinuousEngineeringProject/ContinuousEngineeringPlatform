package main

import (
	"github.com/ContinuousEngineeringProject/cePlatform/docker/bash"
)

func main() {
	manager := bash.ArgsVirtualbox{"manager","virtualbox", "4096","2"}
	bash.RunCmd(bash.CreateVM(manager))

}