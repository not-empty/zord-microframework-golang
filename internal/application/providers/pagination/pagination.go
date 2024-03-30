package pagination

import (
	"go-skeleton/internal/application/services"
	"go-skeleton/internal/repositories/base_repository"
	"math"
	"net/http"
)

type IPaginationProvider[Row any] interface {
	PaginationHandler(page int, limit int) (*services.Error, *Pagination[Row])
}

type Pagination[Row any] struct {
	CurrentPage int
	TotalPages  int64
	Data        *[]Row
}

type PaginationProvider[Row base_repository.Domain] struct {
	repo base_repository.BaseRepository[Row]
}

func NewPaginationProvider[Row base_repository.Domain](repo base_repository.BaseRepository[Row]) *PaginationProvider[Row] {
	return &PaginationProvider[Row]{
		repo: repo,
	}
}

func (pp *PaginationProvider[Row]) PaginationHandler(page int, limit int) (*services.Error, *Pagination[Row]) {
	listData := &[]Row{}
	offset := (page - 1) * limit

	total, err := pp.repo.Count()
	if err != nil {
		return &services.Error{
			Status:  http.StatusInternalServerError,
			Message: "error on pagination count",
		}, nil
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
		listData, err = pp.repo.List(limit, offset)
		if err != nil {
			return &services.Error{
				Status:  http.StatusInternalServerError,
				Message: "error on data list",
			}, nil
		}
	}

	return nil, &Pagination[Row]{
		CurrentPage: page,
		TotalPages:  int64(totalPages),
		Data:        listData,
	}
}
