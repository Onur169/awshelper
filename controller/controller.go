package controller

type Source func() string

type ContentReceiver interface {
	Receive() string
}

type Ctrl struct{}
