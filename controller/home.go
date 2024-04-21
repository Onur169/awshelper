package controller

import (
	"onursahin.dev/awshelper/util"
)

const FileCouldSaveSuccessfully = "file could be saved successfully"

func (c *Ctrl) Home(content func() string) func() {
	return func() {
		c.IsLoadingChannel <- true

		err := util.WriteIntoAwsCredentials(content())
		if err != nil {
			c.HomeChannel <- err.Error()
			c.IsLoadingChannel <- false
			return
		}
		c.HomeChannel <- FileCouldSaveSuccessfully
		c.IsLoadingChannel <- false
	}
}
