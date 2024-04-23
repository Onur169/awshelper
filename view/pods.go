package view

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"onursahin.dev/awshelper/controller"
	"onursahin.dev/awshelper/util"
)

func Pods(c *controller.Ctrl, pods []util.Pod) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ein Fehler ist aufgetreten:", r)
		}
	}()

	podApp := app.New()
	podWindow := podApp.NewWindow("Pods")
	podWindow.Resize(fyne.NewSize(util.AppWidth, util.AppHeight))

	var data = pods
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i].Name)
		},
	)

	podWindow.SetContent(list)
	podWindow.Show()
}
