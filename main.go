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
		HomeChannel:    make(chan string),
		ActionsChannel: make(chan string),
		Tab:            "home",
	}

	myApp := app.New()
	appWindow := myApp.NewWindow("awshelper")
	appWindow.Resize(fyne.NewSize(750, 250))

	tabs := container.NewAppTabs(
		container.NewTabItem("home", view.Home(c)),
		container.NewTabItem("actions", view.Actions(c)),
	)

	tabs.OnSelected = func(item *container.TabItem) {
		c.Tab = item.Text
	}

	tabs.SetTabLocation(container.TabLocationTop)

	appWindow.SetContent(tabs)
	appWindow.ShowAndRun()
}
