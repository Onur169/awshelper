package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"onursahin.dev/awshelper/controller"
	"onursahin.dev/awshelper/view"
)

func main() {
	c := &controller.Ctrl{}

	myApp := app.New()
	appWindow := myApp.NewWindow("awshelper")

	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", view.Home(c)),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	appWindow.SetContent(tabs)
	appWindow.ShowAndRun()
}
