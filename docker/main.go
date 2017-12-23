package main

import (

	"github.com/ContinuousEngineeringProject/cePlatform/docker/bash"
)

// manager nodes
const mName  =  "manager"
const mDriver = "virtualbox"
const mMemory = "1024"
const mCPU  = "1"
const mCount  = 1

// worker nodes
const wName  =  "worker"
const wDriver = "virtualbox"
const wMemory = "1024"
const wCPU  = "1"
const wCount  = 1

func main() {

	// Create manager nodes
	manager := bash.ArgsCreateVM{mName,mDriver, mMemory,mCPU, mCount}
	bash.CreateVM(manager)

	// Create worker nodes
	worker := bash.ArgsCreateVM{wName,wDriver, wMemory,wCPU, wCount}
	bash.CreateVM(worker)

}