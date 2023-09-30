package dummy

import domain "go-skeleton/application/domain/dummy"

type Response struct {
	Status int             `json:"status,omitempty"`
	Data   *[]domain.Dummy `json:"data,omitempty"`
}
