package main

import (
	"github.com/cert-manager/cert-manager/pkg/acme/webhook/cmd"
	"github.com/moritzj29/cert-manager-ddnss-de-webhook/src/ddnss"
	"os"
)

var GroupName = os.Getenv("GROUP_NAME")

func main() {
	if GroupName == "" {
		panic("GROUP_NAME must be specified")
	}

	cmd.RunWebhookServer(GroupName,
		ddnss.NewSolver(nil),
	)
}
