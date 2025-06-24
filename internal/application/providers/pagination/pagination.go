package pagination

import (
	"errors"
	"go-skeleton/internal/application/providers/filters"
	"math"
)

type IPaginationProvider[Row any] interface {
	PaginationHandler(row Row, page int, limit int) (error, *Pagination[Row])
}

type Pagination[Row any] struct {
	CurrentPage int
	TotalPages  int64
	Data        *[]Row
}

type Domain interface {
	Schema() string
	GetFilters() filters.Filters
	SoftDelete() string
}

type IPaginationRepository[Row any] interface {
	List(domain Row, limit int, offset int) (*[]Row, error)
	Count(domain Row) (int64, error)
}

type Provider[Row Domain] struct {
	repo IPaginationRepository[Row]
}

func NewPaginationProvider[Row Domain](repo IPaginationRepository[Row]) *Provider[Row] {
	return &Provider[Row]{
		repo: repo,
	}
}

func (pp *Provider[Row]) PaginationHandler(row Row, page int, limit int) (error, *Pagination[Row]) {
	listData := &[]Row{}
	offset := (page - 1) * limit
	total, err := pp.repo.Count(row)
	if err != nil {
		return errors.New("error on pagination count"), nil
	}

	if total == 0 {
		return nil, &Pagination[Row]{
			CurrentPage: page,
			TotalPages:  0,
			Data:        listData,
		}
	}

	totalPages := math.Ceil(float64(total) / float64(limit))
	if page <= int(totalPages) {
		listData, err = pp.repo.List(row, limit, offset)
		if err != nil {
			return errors.New("error on data list"), nil
		}
	}

	return nil, &Pagination[Row]{
		CurrentPage: page,
		TotalPages:  int64(totalPages),
		Data:        listData,
	}
}
