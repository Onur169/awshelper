package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/url"
	"onursahin.dev/awshelper/controller"
)

func About(c *controller.Ctrl) *fyne.Container {
	label := widget.NewLabel("Hallo!")
	label2 := widget.NewLabel("Das Repository findest du auf: ")

	hyperlink := widget.NewHyperlink("https://github.com/Onur169/awshelper", &url.URL{
		Scheme: "https",
		Host:   "github.com",
		Path:   "Onur169/awshelper",
	})

	return controller.ActionsWrapper(container.NewVBox(
		label,
		label2,
		hyperlink,
	), c)
}
