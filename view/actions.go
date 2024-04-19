package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"onursahin.dev/awshelper/controller"
)

// aws ecr get-login-password --region eu-central-1 | docker login --username AWS --password-stdin 175218586454.dkr.ecr.eu-central-1.amazonaws.com
// kubectl get pods --namespace ma4b
// kubectl logs -n ma4b ma4b-lugas-safe-server-deployment-...-...

const AwsLoginCmd = "aws ecr get-login-password --region eu-central-1 | docker login --username AWS --password-stdin 175218586454.dkr.ecr.eu-central-1.amazonaws.com"
const KubectlGetPodsCmd = "kubectl get pods --namespace ma4b"
const LsLaCmd = "ls -la"
const EchoCmd = "echo $HOME"

func Actions(c *controller.Ctrl) *fyne.Container {
	label := widget.NewLabel("Wähle ein Command aus: ")
	radio := widget.NewRadioGroup([]string{LsLaCmd, EchoCmd}, c.Actions())

	return controller.App(container.NewVBox(
		label,
		radio,
	), c)
}
