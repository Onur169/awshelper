# Awshelper

This project is built using Go programming language version 1.22 with Go SDK 1.22.2.

## Getting Started

These instructions describe how to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go version 1.22: You can download it from the [official Go website](https://golang.org/dl/).
- Set up GOPATH. GOPATH is an environment variable which lists places to look for Go code. Here is a [guide on how to set the GOPATH](https://golang.org/doc/gopath_code.html).

### Installing

1. Clone the repository.
2. Navigate to the project directory.
3. Run `go get .` to download the necessary dependencies.
4. Run `go build` to compile the project.

### Running the Application

After successfully building the project, you can run it by using following command:

`go run .`

## Environment Variables

This project makes use of several environment variables. Make sure to set these variables in your system or in the configuration of your IDE:

- `MOCK_PODS`: This environment variable mock pod list instead of calling `kubectl-get-pods`

## Dependencies

This project uses a number of Go packages:

- fyne.io/fyne/v2 v2.4.5: A modern, easy to use, feature-rich, and async-ready Go web framework and web application deployment tool.
- fyne.io/systray v1.10.1-0.20231115130155-104f5ef7839e: Allows Go applications to create and interact with the system tray on Windows, Linux, and MacOS.
- github.com/joho/godotenv v1.5.1: A Go port of Ruby’s dotenv library (Loads environment variables from .env).
- github.com/stretchr/testify v1.9.0: A toolkit with common assertions and mocks that plays nicely with the standard library.
- github.com/yuin/goldmark v1.7.1: A markdown parser and renderer in Go. It is fully compliant with the CommonMark spec.
- golang.org/x/image v0.15.0: Additional imaging packages for Go.

...and many other indirect dependencies. For a full list of the dependencies, please check the `go.mod` file.
### Author

Onur Sahin