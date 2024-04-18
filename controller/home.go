package controller

import "onursahin.dev/awshelper/util"

func (c *Ctrl) Home(rec ContentReceiver) func() {
	return func() {
		err := util.WriteIntoAwsCredentials(rec.Receive())
		if err != nil {
			println("Fehler beim Schreiben der Datei!", err.Error())
			return
		}
		println("Datei konnte erfolgreich beschrieben werden!")
	}
}
