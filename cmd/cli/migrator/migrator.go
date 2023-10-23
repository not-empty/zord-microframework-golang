package migrator

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-skeleton/pkg"
	"go-skeleton/pkg/database"
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
	migratorInstance := migrator.NewMigrator(pkg.MigratorDependencies["mysql"].(*database.MySql))
	migratorInstance.MigrateAllDomains()
}

func (m *Migrator) BootMigrator(_ *cobra.Command, _ []string) {
	for index, dep := range pkg.MigratorDependencies {
		dep.Boot()
		pkg.Logger.Info(fmt.Sprintf("[Kernel.Kernel] Booting %s", index))
	}
}
