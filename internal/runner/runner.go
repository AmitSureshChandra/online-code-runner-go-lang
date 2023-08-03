package runner

import (
	"online-compiler/internal/docker"
	"online-compiler/internal/file"
	"online-compiler/internal/shell"
	"online-compiler/internal/utils"
	"os/exec"
	"time"
)

func Run(code string, input string, dockerContainerName string, compiler string) ([]byte, time.Time, time.Time) {
	rootFolder := "temp-user" + "/" + utils.GenerateRandomString(10)

	file.CreateFolderIfNotExists(rootFolder)

	file.WriteFile(rootFolder+"/"+docker.GetFileName(compiler), code)
	file.WriteFile(rootFolder+"/input.txt", input)

	containerName := utils.GenerateRandomString(10)
	shell.GetOutput(shell.GetRunCommand(containerName, dockerContainerName, rootFolder))

	start := time.Now()
	time.Sleep(1 * time.Second)
	shell.GetOutput(exec.Command("docker", "stop", containerName))
	stop := time.Now()
	output := shell.GetOutput(exec.Command("docker", "logs", containerName))
	//shell.GetOutput(exec.Command("docker", "rm", containerName))
	shell.GetOutput(exec.Command("rm", "-rf", rootFolder))
	return output, start, stop
}
