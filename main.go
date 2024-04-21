package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"onursahin.dev/awshelper/controller"
	"onursahin.dev/awshelper/util"
	"onursahin.dev/awshelper/view"
)

func main() {
	c := &controller.Ctrl{
		HomeChannel:      make(chan string),
		ActionsChannel:   make(chan string),
		IsLoadingChannel: make(chan bool),
		ResultLabel:      widget.NewLabel(""),
		LoadingLabel:     widget.NewLabel(""),
	}

	myApp := app.New()
	appWindow := myApp.NewWindow("awshelper")
	appWindow.Resize(fyne.NewSize(util.AppWidth, 0))

	homeView := view.Home(c)
	actionsView := view.Actions(c)

	tabList := []*container.TabItem{
		container.NewTabItem("home", homeView),
		container.NewTabItem("actions", actionsView),
	}

	tabs := container.NewAppTabs(tabList...)
	tabs.OnSelected = func(t *container.TabItem) {
		c.ResultLabel.Text = ""
		c.LoadingLabel.Text = ""
		c.ResultLabel.Refresh()
		c.LoadingLabel.Refresh()

		/*		if t.Text == "actions" {
				homeView.Move(fyne.NewPos(-1, -1))
			}*/
	}
	tabs.SetTabLocation(container.TabLocationTop)

	appWindow.SetContent(tabs)
	appWindow.ShowAndRun()
}
