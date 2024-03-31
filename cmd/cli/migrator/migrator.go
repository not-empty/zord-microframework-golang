package migrator

import (
	"fmt"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"
	"go-skeleton/tools/migrator"
	"net/url"

	"github.com/spf13/cobra"
)

type Migrator struct {
	dsn     string
	dsnTest string
}

func NewMigrator() *Migrator {
	return &Migrator{}
}

func (m *Migrator) DeclareCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		&cobra.Command{
			Use:    "migrate",
			Short:  "migrate HCL description database",
			PreRun: m.BootMigrator,
			Run:    m.Migrate,
		},
		&cobra.Command{
			Use:    "inspect",
			Short:  "inspect HCL description database",
			PreRun: m.BootMigrator,
			Run:    m.Inspect,
		},
		&cobra.Command{
			Use:    "generate",
			Short:  "generate HCL from database",
			PreRun: m.BootMigrator,
			Run:    m.Generate,
		},
		&cobra.Command{
			Use:    "migrate-from-domain",
			Short:  "generate HCL from database",
			PreRun: m.BootMigrator,
			Run:    m.MigrateFromDomains,
		},
	)
}

func (m *Migrator) Migrate(_ *cobra.Command, _ []string) {
	migratorInstance := migrator.NewMigrator(m.dsn, m.dsnTest)
	migratorInstance.MigrateAllDomains()
}

func (m *Migrator) MigrateFromDomains(_ *cobra.Command, _ []string) {
	migratorInstance := migrator.NewMigrator(m.dsn, m.dsnTest)
	migratorInstance.MigrateFromDomains()
}

func (m *Migrator) Inspect(_ *cobra.Command, _ []string) {
	migratorInstance := migrator.NewMigrator(m.dsn, m.dsnTest)
	migratorInstance.Inspect()
}

func (m *Migrator) Generate(_ *cobra.Command, _ []string) {
	migratorInstance := migrator.NewMigrator(m.dsn, m.dsnTest)
	migratorInstance.Generate()
}

func (m *Migrator) BootMigrator(_ *cobra.Command, _ []string) {
	conf := config.NewConfig()
	err := conf.LoadEnvs()
	if err != nil {
		panic(err)
	}

	dsn := "%s:%s@%s:%s/%s"
	m.dsn = fmt.Sprintf(
		dsn,
		url.QueryEscape(conf.ReadConfig("DB_USER")),
		url.QueryEscape(conf.ReadConfig("DB_PASS")),
		url.QueryEscape(conf.ReadConfig("DB_URL")),
		conf.ReadConfig("DB_PORT"),
		conf.ReadConfig("DB_DATABASE"),
	)

	m.dsnTest = fmt.Sprintf(
		dsn,
		conf.ReadConfig("DB_USER"),
		conf.ReadConfig("DB_PASS"),
		conf.ReadConfig("DB_URL"),
		conf.ReadConfig("DB_PORT"),
		conf.ReadConfig("DB_TEST_DATABASE"),
	)

	l := logger.NewLogger(
		conf.ReadConfig("ENVIRONMENT"),
		conf.ReadConfig("APP"),
		conf.ReadConfig("VERSION"),
	)

	l.Boot()
}
