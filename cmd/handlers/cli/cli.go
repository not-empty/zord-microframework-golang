package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-skeleton/pkg"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"
)

type Cli struct {
	Environment string
	config      *config.Config
	logger      *logger.Logger
}

func NewCli(Environment string) *Cli {
	c := pkg.KernelDependencies["config"]
	l := pkg.KernelDependencies["logger"]

	return &Cli{
		Environment: Environment,
		config:      c.(*config.Config),
		logger:      l.(*logger.Logger),
	}
}

func (c *Cli) RegisterCommands(cmd *cobra.Command) {

	cmd.AddCommand(
		&cobra.Command{
			Use:    "create-domain",
			Short:  "Create a new domain service",
			Run:    c.BootCli,
			PreRun: c.CreateDomain,
		},
	)

}

func (c *Cli) CreateDomain(cmd *cobra.Command, args []string) {

}

func (c *Cli) BootCli(cmd *cobra.Command, args []string) {
	for index, dep := range pkg.CliDependencies {
		dep.Boot()
		c.logger.Info(fmt.Sprintf("[cli.cli] Booting %s", index))
	}
}
