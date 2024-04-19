package controller

import (
	"errors"
	"onursahin.dev/awshelper/util"
)

var FileCouldNotSavedErr = errors.New("file could be saved successfully")

func (c *Ctrl) Home(rec ContentReceiver) func() {
	return func() {
		err := util.WriteIntoAwsCredentials(rec.Receive())
		if err != nil {
			c.EventChannel <- err.Error()
			return
		}
		c.EventChannel <- FileCouldNotSavedErr.Error()
	}
}
