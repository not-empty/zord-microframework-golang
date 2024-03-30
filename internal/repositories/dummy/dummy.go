package dummyRepository

import (
	"github.com/jmoiron/sqlx"
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/repositories/base_repository"
)

type DummyRepository struct {
	*base_repository.BaseRepo[dummy.Dummy]
}

func NewDummyRepo(mysql *sqlx.DB) *DummyRepository {
	return &DummyRepository{
		BaseRepo: base_repository.NewBaseRepository[dummy.Dummy](mysql),
	}
}
