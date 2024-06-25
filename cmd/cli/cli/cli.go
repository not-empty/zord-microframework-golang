package cli

import (
	"go-skeleton/cmd/cli/domain_reader"
	"go-skeleton/cmd/cli/generator"
	"go-skeleton/cmd/cli/migrator"

	"github.com/spf13/cobra"
)

type Cli struct {
	Cmd *cobra.Command
}

func NewCli(cmd *cobra.Command) *Cli {
	return &Cli{
		Cmd: cmd,
	}
}

func (c *Cli) Start() {
	generatorInstance := generator.NewGenerator()
	generatorInstance.DeclareCommands(c.Cmd)
	migratorInstance := migrator.NewMigrator()
	migratorInstance.DeclareCommands(c.Cmd)
	reader := domain_reader.NewReader()
	reader.DeclareCommands(c.Cmd)
}
