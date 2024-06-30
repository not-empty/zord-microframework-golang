package cli

import (
	"go-skeleton/cmd/cli/generator"
	"go-skeleton/cmd/cli/migrator"
	"go-skeleton/cmd/cli/struct_reader"
	"go-skeleton/pkg/config"

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
	config := config.NewConfig()
	reader := struct_reader.NewStructReader(config)
	reader.DeclareCommands(c.Cmd)
}
