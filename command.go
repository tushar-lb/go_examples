package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	packageManagerExecutable = "kubectl"
)

// main start a gRPC server and waits for connection
func main() {
	cmdStr := "kubectl --kubeconfig px3.yaml apply -f service_account.yaml -n test"
	logrus.Printf("Command : [%s]", cmdStr)
	output, error := executeCmd(cmdStr, true)
	logrus.Printf("Output:[%s], Error:[%+v]", output, error)
}

func executeCmd(cmdStr string, withOutput bool) ([]byte, error) {
	cmd := exec.Command(packageManagerExecutable, cmdStr)
	currEnv := os.Environ()
	cmd.Env = currEnv
	cmd.Args = strings.Fields(cmdStr)
	if !withOutput {
		if err := cmd.Run(); err != nil {
			logrus.Errorf("Failed to %s the service: %+v", cmdStr, err)
			return nil, err
		}
	} else if withOutput {
		output, err := cmd.CombinedOutput()
		if err != nil {
			logrus.Errorf("Failed to %s the service: %+v, %+v", cmdStr, err, string(output))
			errStr := fmt.Sprintf("%s: %s", err.Error(), string(output))
			return nil, errors.New(errStr)
		}
		return output, nil
	}
	return nil, nil
}
