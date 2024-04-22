package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
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
	appWindow.Resize(fyne.NewSize(util.AppWidth, util.AppHeight))

	viewWrapper := func(content *fyne.Container) *fyne.Container {
		return container.NewBorder(
			layout.NewSpacer(),
			layout.NewSpacer(),
			layout.NewSpacer(),
			layout.NewSpacer(),
			content,
		)
	}
	homeView := viewWrapper(view.Home(c))
	actionsView := viewWrapper(view.Actions(c))

	homeTab := container.NewTabItemWithIcon("home", theme.HomeIcon(), homeView)
	actionsTab := container.NewTabItemWithIcon("actions", theme.ComputerIcon(), actionsView)

	tabList := []*container.TabItem{
		homeTab,
		actionsTab,
	}

	tabs := container.NewAppTabs(tabList...)
	tabs.OnSelected = func(t *container.TabItem) {
		c.ResultLabel.Text = ""
		c.LoadingLabel.Text = ""
		c.ResultLabel.Refresh()
		c.LoadingLabel.Refresh()
	}
	tabs.SetTabLocation(container.TabLocationTop)

	appWindow.SetContent(tabs)
	appWindow.ShowAndRun()
}
