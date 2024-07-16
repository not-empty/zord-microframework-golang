package dummy

import (
	"go-skeleton/internal/application/providers/pagination"
	"go-skeleton/internal/repositories/base_repository"
)

type Dummy struct {
	ID        string `db:"id"`
	DummyName string `db:"name"`
	Email     string `db:"email"`
	client    string
}

func (d Dummy) SetClient(client string) Dummy {
	d.client = client
	return d
}

func (d Dummy) Schema() string {
	if d.client == "" {
		return "dummy"
	}
	return d.client + "." + "dummy"
}

type Repository interface {
	base_repository.BaseRepository[Dummy]
}

type PaginationProvider interface {
	pagination.IPaginationProvider[Dummy]
}
