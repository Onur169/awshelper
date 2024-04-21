package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"onursahin.dev/awshelper/controller"
	"onursahin.dev/awshelper/util"
)

func Actions(c *controller.Ctrl) *fyne.Container {
	var cmdList []string
	for k := range util.CommandMap() {
		cmdList = append(cmdList, k)
	}

	radio := widget.NewRadioGroup(cmdList, c.Actions())

	return controller.ActionsWrapper(container.NewVBox(
		radio,
	), c)
}
