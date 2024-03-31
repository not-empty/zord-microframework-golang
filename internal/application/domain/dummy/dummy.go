package dummy

import (
	"go-skeleton/internal/application/providers/pagination"
	"go-skeleton/internal/repositories/base_repository"
)

type Dummy struct {
	ID        string `db:"id" zord_db:"type=char(26),null=false,PK"`
	DummyName string `db:"name" zord_db:"type=char(255),null=false"`
	Email     string `db:"email" zord_db:"type=char(255),null=false,IDX"`
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
