package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"onursahin.dev/awshelper/controller"
	"onursahin.dev/awshelper/util"
)

func Actions(c *controller.Ctrl) *fyne.Container {
	label := widget.NewLabel("Wähle ein Command aus: ")

	radio := widget.NewRadioGroup([]string{
		util.AwsLoginCmd,
		util.KubectlGetPodsCmd,
		util.SleepCmd,
		util.LsCmd,
	}, c.Actions())

	return controller.ActionsWrapper(container.NewVBox(
		label,
		radio,
	), c)
}
