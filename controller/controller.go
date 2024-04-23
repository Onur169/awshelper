package controller

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"onursahin.dev/awshelper/util"
)

type ActionsChannelMsg struct {
	TriggeredCmdKey string
	TriggeredCmd    string
	CmdOutput       string
}

type Ctrl struct {
	HomeChannel      chan string
	ActionsChannel   chan ActionsChannelMsg
	PodsChannel      chan []util.Pod
	IsLoadingChannel chan bool
	ResultLabel      *widget.Label
	LoadingLabel     *widget.Label
}

func HomeWrapper(content *fyne.Container, c *Ctrl) *fyne.Container {
	go func() {
		for {
			select {
			case homeMsg := <-c.HomeChannel:
				c.ResultLabel.Text = homeMsg
				c.ResultLabel.Refresh()
			case isLoadingMsg := <-c.IsLoadingChannel:
				c.LoadingLabel.Text = util.IsLoadingMsg(isLoadingMsg)
				c.LoadingLabel.Refresh()
			}
		}
	}()
	return wrapperBoxV(content, c)
}

func ActionsWrapper(content *fyne.Container, c *Ctrl) *fyne.Container {
	go func() {
		for {
			select {
			case actionsMsg := <-c.ActionsChannel:
				c.ResultLabel.Text = actionsMsg.CmdOutput
				c.ResultLabel.Refresh()

				if actionsMsg.TriggeredCmdKey == "kubectl-get-pods" {
					if util.GetMockPodsEnv() {
						c.PodsChannel <- util.MockPods()
					} else if pods, err := util.ParsePods(actionsMsg.CmdOutput); err == nil {
						c.PodsChannel <- pods
					}
				}
			case isLoadingMsg := <-c.IsLoadingChannel:
				c.LoadingLabel.Text = util.IsLoadingMsg(isLoadingMsg)
				c.LoadingLabel.Refresh()
			}
		}
	}()
	return wrapperBoxH(content, c)
}

func wrapperBoxH(content *fyne.Container, c *Ctrl) *fyne.Container {
	statusBoxContainer := container.NewVBox(c.LoadingLabel, c.ResultLabel)
	statusBoxScroll := container.NewScroll(statusBoxContainer)
	statusBoxScroll.SetMinSize(fyne.NewSize(util.AppWidth, util.AppHeight))
	statusBoxScroll.Refresh()

	contentScroll := container.NewVScroll(content)
	contentScroll.SetMinSize(fyne.NewSize(util.AppWidth/3, util.AppHeight))
	contentScroll.Refresh()

	return container.NewHBox(
		contentScroll,
		statusBoxScroll,
	)
}

func wrapperBoxV(content *fyne.Container, c *Ctrl) *fyne.Container {
	statusBoxContainer := container.NewVBox(c.LoadingLabel, c.ResultLabel)
	statusBoxScroll := container.NewVScroll(statusBoxContainer)

	statusBoxScroll.SetMinSize(fyne.NewSize(util.AppWidth, util.AppHeight))
	statusBoxScroll.Refresh()

	return container.NewBorder(
		nil,
		statusBoxScroll,
		nil,
		nil,
		content,
	)
}
