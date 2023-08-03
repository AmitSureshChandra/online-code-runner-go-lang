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
	newFolder := "temp-user"

	file.CreateFolderIfNotExists(newFolder)

	file.WriteFile(newFolder+"/"+docker.GetFileName(compiler), code)
	file.WriteFile(newFolder+"/input.txt", input)

	containerName := utils.GenerateRandomString(10)
	folderName := utils.GenerateRandomString(10)
	shell.GetOutput(shell.GetRunCommand(containerName, dockerContainerName, folderName))

	start := time.Now()
	time.Sleep(1 * time.Second)
	shell.GetOutput(exec.Command("docker", "stop", containerName))
	stop := time.Now()
	output := shell.GetOutput(exec.Command("docker", "logs", containerName))
	//shell.GetOutput(exec.Command("docker", "rm", containerName))
	shell.GetOutput(exec.Command("rm", "-rf", "temp-user"))
	return output, start, stop
}
