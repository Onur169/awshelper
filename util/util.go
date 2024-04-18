package util

import (
	"os"
)

func WriteIntoAwsCredentials(content string) error {
	filePath := os.ExpandEnv("$HOME/.aws/credentials")

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write([]byte(content)); err != nil {
		return err
	}

	return nil
}
