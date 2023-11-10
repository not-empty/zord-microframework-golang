package migrator

import (
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/logger"
	"go-skeleton/tools/migrator"

	"github.com/spf13/cobra"
)

type Migrator struct {
	db *database.MySql
}

func NewMigrator() *Migrator {
	return &Migrator{}
}

func (m *Migrator) DeclareCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		&cobra.Command{
			Use:    "migrate",
			Short:  "Migrate Gorm Database",
			PreRun: m.BootMigrator,
			Run:    m.Migrate,
		},
	)
}

func (m *Migrator) Migrate(_ *cobra.Command, _ []string) {
	migratorInstance := migrator.NewMigrator(m.db)
	migratorInstance.MigrateAllDomains()
}

func (m *Migrator) BootMigrator(_ *cobra.Command, _ []string) {
	conf := config.NewConfig()

	l := logger.NewLogger(
		conf.ReadConfig("ENVIRONMENT"),
		conf.ReadConfig("APP"),
		conf.ReadConfig("VERSION"),
	)

	l.Boot()

	db := database.NewMysql(
		l,
		conf.ReadConfig("DB_USER"),
		conf.ReadConfig("DB_PASS"),
		conf.ReadConfig("DB_URL"),
		conf.ReadConfig("DB_PORT"),
		conf.ReadConfig("DB_DATABASE"),
	)

	db.Connect()

	m.db = db
}
