package controller

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Source func() string

type ContentReceiver interface {
	Receive() string
}

type Ctrl struct {
	EventChannel chan string
}

func App(content *fyne.Container, c *Ctrl) *fyne.Container {
	statusLabelTxt := func(msg string) string {
		return "Status: " + msg
	}
	statusLabel := widget.NewLabel(statusLabelTxt(""))

	go func() {
		for {
			select {
			case msg := <-c.EventChannel:
				println(msg)
				statusLabel.Text = statusLabelTxt(msg)
				statusLabel.Refresh()
			}
		}
	}()

	return container.NewVBox(layout.NewSpacer(), content, statusLabel)
}
