package util

import (
	"os"
	"os/exec"
)

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
