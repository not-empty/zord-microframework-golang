package migrator

import (
	"go-skeleton/application/domain/dummy"

	//{{codeGen1}}
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
	//{{codeGen2}}
}
