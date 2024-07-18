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
	dsn      string
	dsnTest  string
	database string
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
			Use:    "generate-schema-from-db",
			Short:  "generate-schema-from-db <schema name>",
			Long:   "generate HCL schema from database connected on env",
			PreRun: m.BootMigrator,
			Run:    m.Generate,
		},
	)
}

func (m *Migrator) Migrate(_ *cobra.Command, args []string) {
	tenant := ""
	if len(args) > 0 {
		tenant = args[0]
	}
	migratorInstance := migrator.NewMigrator(m.dsn, m.dsnTest, m.database)
	migratorInstance.MigrateAllDomains(tenant)
}

func (m *Migrator) Inspect(_ *cobra.Command, _ []string) {
	migratorInstance := migrator.NewMigrator(m.dsn, m.dsnTest, m.database)
	migratorInstance.Inspect()
}

func (m *Migrator) Generate(_ *cobra.Command, args []string) {
	schema := ""
	if len(args) > 1 {
		schema = args[1]
	}
	migratorInstance := migrator.NewMigrator(m.dsn, m.dsnTest, m.database)
	migratorInstance.Generate(schema)
}

func (m *Migrator) BootMigrator(_ *cobra.Command, _ []string) {
	conf := config.NewConfig()
	err := conf.LoadEnvs()
	if err != nil {
		panic(err)
	}

	m.database = conf.ReadConfig("DB_DATABASE")
	dsn := "%s:%s@%s:%s"
	m.dsn = fmt.Sprintf(
		dsn,
		url.QueryEscape(conf.ReadConfig("DB_USER")),
		url.QueryEscape(conf.ReadConfig("DB_PASS")),
		url.QueryEscape(conf.ReadConfig("DB_URL")),
		conf.ReadConfig("DB_PORT"),
	)

	testDsn := "%s:%s@%s:%s/%s"
	m.dsnTest = fmt.Sprintf(
		testDsn,
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
