package controller

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strings"
)

type Source func() string

type ContentReceiver interface {
	Receive() string
}

type Ctrl struct {
	Tab            string
	HomeChannel    chan string
	ActionsChannel chan string
}

func HomeWrapper(content *fyne.Container, statusLabel *widget.Label, c *Ctrl) *fyne.Container {
	go func() {
		for {
			select {
			case homeMsg := <-c.HomeChannel:
				statusLabel.Text = homeMsg
				statusLabel.Refresh()
			}
		}
	}()
	return container.NewVBox(layout.NewSpacer(), content, statusLabel)
}

func ActionsWrapper(content *fyne.Container, statusLabel *widget.Label, c *Ctrl) *fyne.Container {
	go func() {
		for {
			select {
			case actionsMsg := <-c.ActionsChannel:
				statusLabel.Text = actionsMsg
				statusLabel.Refresh()
			}
		}
	}()
	return container.NewVBox(layout.NewSpacer(), content, statusLabel)
}
