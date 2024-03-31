package migrator

import (
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/repositories/base_repository"
)

func GetTables() []base_repository.Domain {
	return []base_repository.Domain{
		new(dummy.Dummy),
		//tables.go
	}
}
