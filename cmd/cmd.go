package cmd

import (
	"github.com/spf13/cobra"
	"go-skeleton/cmd/handlers/cli"
	"go-skeleton/cmd/handlers/http"
)

type Command interface {
	BaseCommand() *cobra.Command
}

var CommandList = map[string]Command{
	"http": http.NewServer(),
	"cli":  cli.NewCli(),
}
