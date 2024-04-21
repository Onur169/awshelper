package controller

import (
	"onursahin.dev/awshelper/util"
)

func (c *Ctrl) Actions() func(string) {
	return func(value string) {
		c.ResultLabel.Text = ""
		c.ResultLabel.Refresh()

		cmdMap := util.CommandMap()
		cmd := cmdMap[value]

		var out string
		var err error
		go func() {
			c.IsLoadingChannel <- true

			out, err = util.RunCommand(cmd)
			if err != nil {
				c.ActionsChannel <- util.CmdErrResult(err)
				c.IsLoadingChannel <- false
				return
			}
			c.ActionsChannel <- util.CmdOutResult(out)

			c.IsLoadingChannel <- false
		}()
	}
}
