package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"onursahin.dev/awshelper/controller"
)

type HomeContentReceiver struct {
	controller.ContentReceiver
	controller.Source
}

func (r *HomeContentReceiver) Receive() string {
	return r.Source()
}

func Home(c *controller.Ctrl) *fyne.Container {
	textArea := widget.NewEntry()
	textArea.MultiLine = true
	textArea.SetMinRowsVisible(10)

	rc := &HomeContentReceiver{Source: func() string {
		return textArea.Text
	}}

	sendButton := widget.NewButton("Absenden", c.Home(rc))
	label := widget.NewLabel("Hier Credentials eingeben:")
	statusLabel := widget.NewLabel(StatusLabelTxt(""))

	return controller.HomeWrapper(container.NewVBox(
		label,
		textArea,
		sendButton,
	), statusLabel, c)
}
