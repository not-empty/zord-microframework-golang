package cli

import (
	"fmt"
	"go-skeleton/pkg"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/logger"
	"go-skeleton/tools/generator"
	"go-skeleton/tools/migrator"

	"github.com/spf13/cobra"
)

var domain string

type Cli struct {
	Environment string
	validator   bool
}

func NewCli() *Cli {
	return &Cli{
		Environment: pkg.Config.ReadConfig("ENVIRONMENT"),
	}
}

func (c *Cli) Start(cmd *cobra.Command) {
	c.initFlags(cmd)
	createDomain := &cobra.Command{
		Use:   "create-domain",
		Short: "Create a new domain service",
		Run:   c.CreateDomain,
	}
	createDomain.Flags().BoolVarP(&c.validator, "validator", "v", false, "Create domain with validation")
	cmd.AddCommand(
		createDomain,
		&cobra.Command{
			Use:   "destroy-domain",
			Short: "Destroy a domain service",
			Run:   c.DestroyDomain,
		},
		&cobra.Command{
			Use:   "migrate",
			Short: "Migrate Gorm Database",
			Run:   c.Migrate,
		},
	)
}

func (c *Cli) CreateDomain(_ *cobra.Command, args []string) {
	generatorInstance := generator.NewGenerator(pkg.CliDependencies["logger"].(*logger.Logger))
	if len(args) > 0 {
		domain = args[0]
	}
	err := generatorInstance.CreateDomain(domain, c.validator)
	pkg.Logger.Error(err)
}

func (c *Cli) DestroyDomain(_ *cobra.Command, args []string) {
	generatorInstance := generator.NewGenerator(pkg.CliDependencies["logger"].(*logger.Logger))
	if len(args) > 0 {
		domain = args[0]
	}
	err := generatorInstance.DestroyDomain(domain)
	pkg.Logger.Error(err)
}

func (c *Cli) Migrate(_ *cobra.Command, _ []string) {
	migratorInstance := migrator.NewMigrator(pkg.MigratorDependencies["mysql"].(*database.MySql))
	migratorInstance.MigrateAllDomains()
}

func (c *Cli) initFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&domain, "domain", "", "Describe name to New Domain")
	cmd.MarkFlagsMutuallyExclusive("domain")
}

func (c *Cli) BootGenerator() {
	for index, dep := range pkg.CliDependencies {
		dep.Boot()
		pkg.Logger.Info(fmt.Sprintf("[Kernel.Kernel] Booting %s", index))
	}
}

func (c *Cli) BootMigrator() {
	for index, dep := range pkg.MigratorDependencies {
		dep.Boot()
		pkg.Logger.Info(fmt.Sprintf("[Kernel.Kernel] Booting %s", index))
	}
}
