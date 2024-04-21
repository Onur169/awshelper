package util

import (
	"os"
	"os/exec"
)

const AppWidth = 750
const AppHeight = 500

func WriteIntoAwsCredentials(content string) error {
	filePath := os.ExpandEnv("$HOME/.aws/credentials")

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	if _, err := file.Write([]byte(content)); err != nil {
		return err
	}

	return nil
}

func RunCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func TruncateString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length] + "..."
}

const AwsLoginCmd = "aws-login"
const KubectlGetPodsCmd = "kubectl-get-pods"
const SleepCmd = "sleep-echo"
const LsCmd = "list files"

func CommandMap() map[string]string {
	m := make(map[string]string)
	m[AwsLoginCmd] = "aws ecr get-login-password --region eu-central-1 | docker login --username AWS --password-stdin 175218586454.dkr.ecr.eu-central-1.amazonaws.com"
	m[KubectlGetPodsCmd] = "kubectl get pods --namespace ma4b"
	m[SleepCmd] = "sleep 4 && echo \"Waited 2 sec\" "
	m[LsCmd] = "ls -la"
	return m
}

func CmdErrResult(err error) string {
	return err.Error()
}

func CmdOutResult(out string) string {
	return out
}

func IsLoadingMsg(isLoading bool) string {
	if isLoading {
		return "Loading..."
	}
	return "Finished!"
}
