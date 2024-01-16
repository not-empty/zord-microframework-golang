package user

import domain "go-skeleton/internal/application/domain/user"

type Response struct {
	CurrentPage int                         `json:"current_page"`
	TotalPages  int64                       `json:"total_pages"`
	Data   *[]domain.User   `json:"data,omitempty"`
}
