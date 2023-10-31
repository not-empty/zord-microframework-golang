package main

import (
	"github.com/spf13/cobra"
	"go-skeleton/cmd/cli/cli"
)

func main() {
	cmd := &cobra.Command{}
	cliInstance := cli.NewCli(cmd)
	cliInstance.Start()
	cliInstance.Cmd.Execute()
}
