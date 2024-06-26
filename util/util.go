package util

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const AppWidth = 500
const AppHeight = 75

type Pod struct {
	Name     string
	Ready    string
	Status   string
	Restarts string
	Age      string
}

func FilterPods(pods []Pod, searchTerm string) []Pod {
	var filteredPods []Pod
	for _, pod := range pods {
		if strings.Contains(pod.Name, searchTerm) {
			filteredPods = append(filteredPods, pod)
		}
	}
	return filteredPods
}

func MockPods() []Pod {
	return []Pod{
		{Name: "Onur-Pod", Ready: "True", Status: "Running", Restarts: "0", Age: "10m"},
		{Name: "Runo-Pod", Ready: "False", Status: "Pending", Restarts: "1", Age: "5m"},
		{Name: "Random-Pod", Ready: "True", Status: "Running", Restarts: "2", Age: "20m"},
		{Name: "Random2-Pod", Ready: "True", Status: "Running", Restarts: "2", Age: "20m"},
		{Name: "Modnar-Pod", Ready: "True", Status: "Running", Restarts: "0", Age: "15m"},
		{Name: "Foobar-Pod", Ready: "False", Status: "Pending", Restarts: "3", Age: "25m"},
		{Name: "Barfoo-Pod", Ready: "False", Status: "Pending", Restarts: "3", Age: "25m"},
	}
}

func ParsePods(input string) ([]Pod, error) {
	var pods []Pod

	lines := strings.SplitN(input, "\n", -1)

	if len(lines) == 0 {
		return nil, errors.New("input is empty")
	}

	if len(lines) < 2 {
		return nil, errors.New("header line is missing")
	}

	lines = lines[1:]

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 5 {
			continue
		}
		pods = append(pods, Pod{
			Name:     fields[0],
			Ready:    fields[1],
			Status:   fields[2],
			Restarts: fields[3],
			Age:      fields[4],
		})
	}

	return pods, nil
}

func ReplaceFirstLine(content string) string {
	lines := strings.Split(content, "\n")

	if len(lines) == 0 {
		return content
	}

	lines[0] = "[default]"

	return strings.Join(lines, "\n")
}

func WriteIntoAwsCredentials(content string) error {
	updatedContent := ReplaceFirstLine(content)

	filePath := os.ExpandEnv("$HOME/.aws/credentials")

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	if _, err := file.Write([]byte(updatedContent)); err != nil {
		return err
	}

	return nil
}

func SetPathForCmdExec() error {
	pathList := []string{
		"/usr/local/bin",
		"/usr/bin",
		"/bin",
		"/usr/sbin",
		"/sbin",
		"/opt/homebrew/bin",
	}
	joinedPath := strings.Join(pathList, ":")
	return os.Setenv("PATH", joinedPath)
}

func RunCommand(command string) (string, error) {
	err := SetPathForCmdExec()
	if err != nil {
		return "", err
	}

	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func CommandMap() map[string]string {
	m := make(map[string]string)
	m["aws-version"] = "aws --version"
	m["kubectl-version"] = "kubectl version"
	m["aws-login"] = "aws ecr get-login-password --region eu-central-1 | docker login --username AWS --password-stdin 175218586454.dkr.ecr.eu-central-1.amazonaws.com"
	m["kubectl-get-pods"] = "kubectl get pods --namespace ma4b"
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

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
}

func GetMockPodsEnv() bool {
	mockPodsStr := os.Getenv("MOCK_PODS")
	if mockPodsStr == "" {
		return false
	}

	mockPods, err := strconv.ParseBool(mockPodsStr)
	if err != nil {
		return false
	}

	return mockPods
}

func OpenCmdWithCommand(command string) {
	err := SetPathForCmdExec()

	cmd := exec.Command(
		"osascript",
		"-e",
		fmt.Sprintf("tell application \"Terminal\" to do script \" %s \"", command),
	)
	err = cmd.Start()
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
		return
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("cmd.Wait() failed with %s\n", err)
		return
	}
}
