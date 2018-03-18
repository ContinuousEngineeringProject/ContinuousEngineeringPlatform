package bash

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func runBashCmd (cmd *exec.Cmd) (status string){
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		status = "Error creating StdoutPipe for Cmd:"
		fmt.Fprintln(os.Stderr, status, err)
//		os.Exit(1)
		return status
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("| %s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		status = "Error starting Cmd:"
		fmt.Fprintln(os.Stderr, status, err)
		return status
	}

	err = cmd.Wait()
	if err != nil {
		status = "Error waiting for Cmd:" + err
		fmt.Fprintln(os.Stderr, status, err)
		return status
	}
	
	return
}
