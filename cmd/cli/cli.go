package main

import (
	"go-skeleton/application/services"
	"go-skeleton/cmd/cli/generator"
	"go-skeleton/cmd/cli/migrator"
	"go-skeleton/pkg"
	"go-skeleton/pkg/logger"

	"github.com/spf13/cobra"
)

type Cli struct {
	Environment string
	cmd         *cobra.Command
	logger      services.Logger
}

func NewCli(cmd *cobra.Command) *Cli {
	l := pkg.CliDependencies["logger"]

	return &Cli{
		Environment: pkg.Config.ReadConfig("ENVIRONMENT"),
		cmd:         cmd,
		logger:      l.(*logger.Logger),
	}
}

func main() {
	cmd := &cobra.Command{}
	cli := NewCli(cmd)
	cli.Start()
	cli.cmd.Execute()
}

func (c *Cli) Start() {
	generatorInstance := generator.NewGenerator(c.logger)
	generatorInstance.DeclareCommands(c.cmd)
	migratorInstance := migrator.NewMigrator()
	migratorInstance.DeclareCommands(c.cmd)
}
