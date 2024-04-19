package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"onursahin.dev/awshelper/controller"
	"onursahin.dev/awshelper/view"
)

func main() {
	c := &controller.Ctrl{
		EventChannel: make(chan string),
	}

	myApp := app.New()
	appWindow := myApp.NewWindow("awshelper")
	appWindow.Resize(fyne.NewSize(750, 250))

	tabs := container.NewAppTabs(
		container.NewTabItem("Home", view.Home(c)),
		container.NewTabItem("AWS Actions", view.Actions(c)),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	appWindow.SetContent(tabs)
	appWindow.ShowAndRun()
}
