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

type PaginationProvider[Row any] struct {
	repo base_repository.BaseRepository[Row]
}

func NewPaginationProvider[Row any](repo base_repository.BaseRepository[Row]) *PaginationProvider[Row] {
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
			Message: "Try again in a few minutes",
			Error:   "fatal error",
		}, nil
	}

	if total == 0 {
		return &services.Error{
			Status:  http.StatusNotFound,
			Message: "Try again in a few minutes",
			Error:   "data not found",
		}, nil
	}

	totalPages := math.Ceil(float64(total) / float64(limit))
	if page <= int(totalPages) {
		listData, err = pp.repo.List(limit, offset)
		if err != nil {
			return &services.Error{
				Status:  http.StatusInternalServerError,
				Message: "Try again in a few minutes",
				Error:   "Error on request process",
			}, nil
		}
	}

	return nil, &Pagination[Row]{
		CurrentPage: page,
		TotalPages:  int64(totalPages),
		Data:        listData,
	}
}
