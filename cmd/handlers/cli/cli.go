package cli

import (
	"fmt"
	"go-skeleton/pkg"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/logger"
	"go-skeleton/tools/generator"
	"go-skeleton/tools/migrator"

	"github.com/spf13/cobra"
)

var domain string

type Cli struct {
	Environment string
	config      *config.Config
	logger      *logger.Logger
	mysql       *database.MySql
}

func NewCli(Environment string) *Cli {
	c := pkg.CliDependencies["config"]
	l := pkg.CliDependencies["logger"]
	m := pkg.CliDependencies["mysql"]

	return &Cli{
		Environment: Environment,
		config:      c.(*config.Config),
		logger:      l.(*logger.Logger),
		mysql:       m.(*database.MySql),
	}
}

func (c *Cli) RegisterCommands(cmd *cobra.Command) {
	c.initFlags(cmd)
	cmd.AddCommand(
		&cobra.Command{
			Use:              "create-domain",
			Short:            "Create a new domain service",
			Run:              c.BootCli,
			PreRun:           c.CreateDomain,
			TraverseChildren: true,
		},
		&cobra.Command{
			Use:    "destroy-domain",
			Short:  "Destroy a domain service",
			Run:    c.BootCli,
			PreRun: c.DestroyDomain,
		},
		&cobra.Command{
			Use:    "migrate",
			Short:  "Migrate Gorm Database",
			Run:    c.Migrate,
			PreRun: c.BootCli,
		},
	)

}

func (c *Cli) CreateDomain(cmd *cobra.Command, args []string) {
	generatorInstance := generator.NewGenerator(c.logger)
	generatorInstance.CreateDomain(domain)
}

func (c *Cli) DestroyDomain(cmd *cobra.Command, args []string) {
	generatorInstance := generator.NewGenerator(c.logger)
	generatorInstance.DestroyDomain(domain)
}

func (c *Cli) Migrate(cmd *cobra.Command, args []string) {
	migratorInstace := migrator.NewMigrator(c.mysql)
	migratorInstace.MigrateAllDomains()
}

func (c *Cli) initFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&domain, "domain", "", "Describe name to New Domain")
	cmd.MarkFlagsMutuallyExclusive("domain")
}

func (c *Cli) BootCli(cmd *cobra.Command, args []string) {
	for index, dep := range pkg.CliDependencies {
		dep.Boot()
		c.logger.Info(fmt.Sprintf("[cli.cli] Booting %s", index))
	}
}
