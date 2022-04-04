package commandUtils

import (
	"os/exec"
)

func ExecCommand(command string) (exitCode int, err error) {
	proc := exec.Command("/bin/bash", "-c", command)
	err = proc.Run()

	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			return
		} else {
			err = nil
		}
	}
	exitCode = proc.ProcessState.ExitCode()
	return
}

func ExecCommandAndGetOutput(command string) (exitCode int, output *string, err error) {
	proc := exec.Command("/bin/bash", "-c", command)
	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			return 1, nil, err
		} else {
			err = nil
		}
	}
	outputBytes, err := proc.CombinedOutput()
	if err != nil {
		return 1, nil, err
	}
	exitCode = proc.ProcessState.ExitCode()
	outputStr := string(outputBytes)
	output = &outputStr

	return
}
