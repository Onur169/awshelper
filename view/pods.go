package view

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"onursahin.dev/awshelper/controller"
	"onursahin.dev/awshelper/util"
)

func buildList(data []util.Pod) *widget.List {
	return widget.NewList(
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
}

func handleOnSelectedListItem(pods []util.Pod) func(id widget.ListItemID) {
	return func(id widget.ListItemID) {
		selectedPod := pods[id]
		cmd := "kubectl logs -n ma4b %s -f"
		go util.OpenCmdWithCommand(fmt.Sprintf(cmd, selectedPod.Name))
	}
}

func buildListContent(searchField *widget.Entry, list *widget.List) *fyne.Container {
	return container.NewBorder(
		searchField,
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		list,
	)
}

func Pods(c *controller.Ctrl, pods []util.Pod) {

	list := buildList(pods)
	list.OnSelected = handleOnSelectedListItem(pods)

	searchField := widget.NewEntry()
	searchField.PlaceHolder = "Nach Pod suchen"
	searchField.OnChanged = func(s string) {
		filteredPods := util.FilterPods(pods, s)
		list = buildList(filteredPods)
		list.OnSelected = handleOnSelectedListItem(filteredPods)
		content := buildListContent(searchField, list)
		c.PodWindow.SetContent(content)
	}

	content := buildListContent(searchField, list)

	c.PodWindow.SetContent(content)
	c.PodWindow.Show()
}
