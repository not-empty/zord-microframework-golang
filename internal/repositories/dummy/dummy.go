package dummyRepository

import (
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/repositories/base_repository"

	"github.com/jmoiron/sqlx"
)

type DummyRepository struct {
	*base_repository.BaseRepo[dummy.Dummy]
}

func NewDummyRepository(mysql *sqlx.DB) *DummyRepository {
	return &DummyRepository{
		BaseRepo: base_repository.NewBaseRepository[dummy.Dummy](mysql),
	}
}
