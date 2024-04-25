package view

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"onursahin.dev/awshelper/controller"
	"onursahin.dev/awshelper/util"
)

func Pods(c *controller.Ctrl, pods []util.Pod) {

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

	list.OnSelected = func(id widget.ListItemID) {
		selectedPod := data[id]
		fmt.Println(selectedPod)
		go util.OpenCmdWithCommand(fmt.Sprintf("kubectl logs -n ma4b %s -f", selectedPod.Name))
	}

	list.BaseWidget.Resize(fyne.NewSize(util.AppWidth, 400))
	list.BaseWidget.Refresh()

	searchField := widget.NewEntry()

	content := container.NewVBox(searchField, list)

	c.PodWindow.SetContent(content)
	c.PodWindow.Show()
}
