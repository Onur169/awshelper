package util

import (
	"os"
	"os/exec"
)

const AppWidth = 500
const AppHeight = 75

func WriteIntoAwsCredentials(content string) error {
	filePath := os.ExpandEnv("$HOME/.aws/credentials")

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
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

func CommandMap() map[string]string {
	m := make(map[string]string)
	m["aws-login"] = "aws ecr get-login-password --region eu-central-1 | docker login --username AWS --password-stdin 175218586454.dkr.ecr.eu-central-1.amazonaws.com"
	m["kubectl-get-pods"] = "kubectl get pods --namespace ma4b"
	m["ls-la-test"] = "ls -la"
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
