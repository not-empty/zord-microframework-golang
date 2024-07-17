package dummy

import (
	"go-skeleton/internal/application/providers/filters"
	"go-skeleton/internal/application/providers/pagination"
	"go-skeleton/internal/repositories/base_repository"
)

type Dummy struct {
	ID        string `db:"id"`
	DummyName string `db:"name"`
	Email     string `db:"email"`
	client    string
	filters   *filters.Filters
}

func (d *Dummy) SetClient(client string) {
	d.client = client
}

func (d *Dummy) SetFilters(filters *filters.Filters) {
	d.filters = filters
}

func (d Dummy) GetFilters() filters.Filters {
	if d.filters != nil {
		return *d.filters
	}
	return filters.Filters{}
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
