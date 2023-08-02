package shell

import (
	"os"
	"os/exec"
)

func GetRunCommand(containerName string, dockerContainerName string) *exec.Cmd {
	curDir, _ := os.Getwd()
	return exec.Command("docker", "run", "--name", containerName, "-d", "--memory", "100mb", "--cpu-quota=100000", "-v", curDir+"/temp-user"+":/opt/myapp", dockerContainerName)
}

func GetOutput(cmd *exec.Cmd) []byte {
	output, _ := cmd.CombinedOutput()
	return output
}
