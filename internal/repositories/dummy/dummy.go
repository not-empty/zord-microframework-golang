package dummyRepository

import (
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/repositories/base_repository"
	"go-skeleton/pkg/database"
)

type DummyRepository struct {
	*base_repository.BaseRepo[dummy.Dummy]
}

func NewDummyRepo(mysql *database.MySql) *DummyRepository {
	return &DummyRepository{
		BaseRepo: base_repository.NewBaseRepository[dummy.Dummy](mysql),
	}
}
