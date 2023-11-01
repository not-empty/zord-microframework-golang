package cli

import (
	"github.com/spf13/cobra"
	"go-skeleton/cmd/cli/generator"
	"go-skeleton/cmd/cli/migrator"
	"go-skeleton/internal/application/services"
	"go-skeleton/pkg"
)

type Cli struct {
	Environment string
	Cmd         *cobra.Command
	logger      services.Logger
}

func NewCli(cmd *cobra.Command) *Cli {

	return &Cli{
		Environment: pkg.Config.ReadConfig("ENVIRONMENT"),
		Cmd:         cmd,
		logger:      pkg.Logger,
	}
}

func (c *Cli) Start() {
	generatorInstance := generator.NewGenerator(c.logger)
	generatorInstance.DeclareCommands(c.Cmd)
	migratorInstance := migrator.NewMigrator()
	migratorInstance.DeclareCommands(c.Cmd)
}
