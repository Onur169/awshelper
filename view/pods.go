package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"onursahin.dev/awshelper/controller"
	"onursahin.dev/awshelper/util"
)

func Pods(c *controller.Ctrl, pods []util.Pod) {

	podApp := app.New()
	podWindow := podApp.NewWindow("Pods")
	podWindow.Resize(fyne.NewSize(util.AppWidth/2, util.AppHeight/2))

	content := container.NewVBox(widget.NewLabel("Hallo Welt"))

	podWindow.SetContent(content)
	podWindow.ShowAndRun()
}
