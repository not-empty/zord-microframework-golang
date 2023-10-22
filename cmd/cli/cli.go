package main

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
	c.initFlags(c.cmd)
	createDomain := &cobra.Command{
		Use:    "create-domain",
		Short:  "Create a new domain service",
		PreRun: c.BootGenerator,
		Run:    c.CreateDomain,
	}
	createDomain.Flags().BoolVarP(&c.validator, "validator", "v", false, "Create domain with validation")
	c.cmd.AddCommand(
		createDomain,
		&cobra.Command{
			Use:    "destroy-domain",
			Short:  "Destroy a domain service",
			PreRun: c.BootGenerator,
			Run:    c.DestroyDomain,
		},
		&cobra.Command{
			Use:    "migrate",
			Short:  "Migrate Gorm Database",
			PreRun: c.BootMigrator,
			Run:    c.Migrate,
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

func (c *Cli) BootGenerator(_ *cobra.Command, _ []string) {
	for index, dep := range pkg.CliDependencies {
		dep.Boot()
		pkg.Logger.Info(fmt.Sprintf("[Kernel.Kernel] Booting %s", index))
	}
}

func (c *Cli) BootMigrator(_ *cobra.Command, _ []string) {
	for index, dep := range pkg.MigratorDependencies {
		dep.Boot()
		pkg.Logger.Info(fmt.Sprintf("[Kernel.Kernel] Booting %s", index))
	}
}
