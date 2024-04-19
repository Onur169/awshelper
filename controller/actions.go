package controller

import (
	"onursahin.dev/awshelper/util"
)

func (c *Ctrl) Actions() func(string) {
	return func(value string) {
		out, err := util.RunCommand(value)
		if err == nil {
			c.ActionsChannel <- out
		}
	}
}
