package migrator

import (
	"go-skeleton/internal/application/domain/dummy"

	"go-skeleton/internal/application/domain/user"
	//{{codeGen3}}
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
	m.db.Db.Migrator().CreateTable(&user.User{})
	//{{codeGen4}}
}
