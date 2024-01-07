package pagination

import (
	"go-skeleton/internal/application/services"
	"go-skeleton/internal/repositories/base_repository"
	"math"
	"net/http"
)

type Pagination[Row any] struct {
	CurrentPage int
	TotalPages  int64
	Data        *[]Row
}

func PaginationProvider[Row any](repo base_repository.BaseRepository[Row], page int, limit int) (*services.Error, *Pagination[Row]) {
	listData := &[]Row{}
	offset := (page - 1) * limit

	total, err := repo.Count()
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
		listData, err = repo.List(limit, offset)
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
