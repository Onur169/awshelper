package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"onursahin.dev/awshelper/controller"
)

func Home(c *controller.Ctrl) *fyne.Container {
	textArea := widget.NewEntry()
	textArea.MultiLine = true
	textArea.SetPlaceHolder("Place your aws credentials here")
	content := func() string { return textArea.Text }

	sendButton := widget.NewButton("Absenden", c.Home(content))

	return controller.HomeWrapper(container.NewVBox(
		textArea,
		sendButton,
	), c)
}
