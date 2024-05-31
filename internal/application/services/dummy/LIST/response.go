package dummy

import (
	domain "go-skeleton/internal/application/domain/dummy"
)

type Response struct {
	CurrentPage int
	TotalPages  int64
	Data        *[]domain.Dummy
}
