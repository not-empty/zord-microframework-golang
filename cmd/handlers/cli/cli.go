package cli

import (
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

func (c *Cli) RegisterCommands(cmd *cobra.Command) {}
