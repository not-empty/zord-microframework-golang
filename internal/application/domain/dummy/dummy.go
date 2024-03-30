package dummy

import (
	"go-skeleton/internal/application/providers/pagination"
	"go-skeleton/internal/repositories/base_repository"
)

type Dummy struct {
	ID        string `db:"id"`
	DummyName string `db:"name"`
}

func (d Dummy) Schema() string {
	return "dummies"
}

type Repository interface {
	base_repository.BaseRepository[Dummy]
}

type PaginationProvider interface {
	pagination.IPaginationProvider[Dummy]
}
