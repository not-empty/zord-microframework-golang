package migrator

import (
	"go-skeleton/application/domain/dummy"
	"go-skeleton/pkg/database"
)

type Migrator struct {
	db *database.MySql
}

func NewMigrator(db *database.MySql) *Migrator {
	return &Migrator{
		db: db,
	}
}

func (m *Migrator) MigrateAllDomains() {
	m.db.Db.Migrator().CreateTable(&dummy.Dummy{})
}
