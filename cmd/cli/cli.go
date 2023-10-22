package main

import (
	"github.com/spf13/cobra"
	"go-skeleton/cmd/cli/generator"
	"go-skeleton/cmd/cli/migrator"
	"go-skeleton/pkg"
)

type Cli struct {
	Environment string
	validator   bool
	cmd         *cobra.Command
}

func NewCli(cmd *cobra.Command) *Cli {
	return &Cli{
		Environment: pkg.Config.ReadConfig("ENVIRONMENT"),
		cmd:         cmd,
	}
}

func main() {
	cmd := &cobra.Command{}
	cli := NewCli(cmd)
	cli.Start()
	cli.cmd.Execute()
}

func (c *Cli) Start() {
	generatorInstance := generator.NewGenerator()
	generatorInstance.DeclareCommand(c.cmd)
	migratorInstance := migrator.NewMigrator()
	migratorInstance.DeclareCommands(c.cmd)
}
