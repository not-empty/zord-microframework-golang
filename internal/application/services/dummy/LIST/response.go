package dummy

import (
	domain "go-skeleton/internal/application/domain/dummy"
)

type Response struct {
	Status      int             `json:"status,omitempty"`
	CurrentPage int             `json:"current_page"`
	TotalPages  int64           `json:"total_pages"`
	Data        *[]domain.Dummy `json:"data,omitempty"`
}
