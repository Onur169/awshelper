package controller

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"onursahin.dev/awshelper/util"

	"fyne.io/fyne/v2/widget"
)

type Ctrl struct {
	HomeChannel      chan string
	ActionsChannel   chan string
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
				c.ResultLabel.Text = actionsMsg
				c.ResultLabel.Refresh()
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

	statusBoxScroll.SetMinSize(fyne.NewSize(util.AppWidth, 100))
	statusBoxScroll.Refresh()

	return container.NewHBox(
		content,
		statusBoxScroll,
	)
}

func wrapperBoxV(content *fyne.Container, c *Ctrl) *fyne.Container {
	statusBoxContainer := container.NewVBox(c.LoadingLabel, c.ResultLabel)

	return container.NewBorder(
		nil,
		statusBoxContainer,
		nil,
		nil,
		content,
	)
}
