package bash

import "testing"

// test data - node
const mName = "manager"
const mDriver = "virtualbox"
const mMemory = "1024"
const mCPU = "1"
const mCount = 1

func TestCreateNode(t *testing.T) {
	// Set test data
	testData := ArgsCreateNode{mName, mDriver, mMemory, mCPU, mCount}

	DmCreate(testData)

	// TODO: Complete DmCreate test
}

func TestRemoveNode(t *testing.T) {
	// Set test data
	testData := ArgsCreateNode{mName, mDriver, mMemory, mCPU, mCount}

	DmRemove(testData.PREFIX + "1")

	// TODO: Complete DmRemove test
}

func TestStartNode(t *testing.T) {
	// Set test data
	testData := ArgsCreateNode{mName, mDriver, mMemory, mCPU, mCount}

	DmStart(testData.PREFIX + "1")

	// TODO: Complete DmStart test
}

func TestStopNode(t *testing.T) {
	// Set test data
	testData := ArgsCreateNode{mName, mDriver, mMemory, mCPU, mCount}

	DmStop(testData.PREFIX + "1")

	// TODO: Complete DmStop test
}
