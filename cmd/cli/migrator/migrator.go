package migrator

import (
	"github.com/spf13/cobra"
	"go-skeleton/pkg"
	"go-skeleton/tools/migrator"
)

type Migrator struct {
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
	migratorInstance := migrator.NewMigrator(pkg.Mysql)
	migratorInstance.MigrateAllDomains()
}

func (m *Migrator) BootMigrator(_ *cobra.Command, _ []string) {
	pkg.Mysql.Connect()
}
