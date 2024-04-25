package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"onursahin.dev/awshelper/controller"
	"onursahin.dev/awshelper/util"
	"onursahin.dev/awshelper/view"
)

func main() {
	myApp := app.New()
	podWindow := myApp.NewWindow("Pods")
	podWindow.Resize(fyne.NewSize(util.AppWidth, util.AppHeight*4))
	podWindow.CenterOnScreen()
	podWindow.SetCloseIntercept(func() {
		podWindow.Hide()
	})

	util.LoadEnv()
	log.Println("MOCK_PODS env =", util.GetMockPodsEnv())

	c := &controller.Ctrl{
		HomeChannel:      make(chan string),
		ActionsChannel:   make(chan controller.ActionsChannelMsg),
		IsLoadingChannel: make(chan bool),
		PodsChannel:      make(chan []util.Pod),
		ResultLabel:      widget.NewLabel(""),
		LoadingLabel:     widget.NewLabel(""),
		PodWindow:        podWindow,
	}

	appWindow := myApp.NewWindow("awshelper")
	appWindow.Resize(fyne.NewSize(util.AppWidth, util.AppHeight))
	appWindow.SetMaster()
	appWindow.CenterOnScreen()

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
	aboutView := viewWrapper(view.About(c))

	homeTab := container.NewTabItemWithIcon("home", theme.HomeIcon(), homeView)
	actionsTab := container.NewTabItemWithIcon("actions", theme.ComputerIcon(), actionsView)
	aboutTab := container.NewTabItemWithIcon("about", theme.InfoIcon(), aboutView)

	tabList := []*container.TabItem{
		homeTab,
		actionsTab,
		aboutTab,
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

	go func() {
		for {
			select {
			case podMsg := <-c.PodsChannel:
				view.Pods(c, podMsg)
			}
		}
	}()

	appWindow.ShowAndRun()
}
