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
	//textArea.SetMinRowsVisible(10)
	content := textArea.Text

	sendButton := widget.NewButton("Absenden", c.Home(content))
	label := widget.NewLabel("Hier Credentials eingeben:")

	return controller.HomeWrapper(container.NewVBox(
		label,
		textArea,
		sendButton,
	), c)
}
