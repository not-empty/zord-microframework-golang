package dummy

import (
	"go-skeleton/internal/application/domain/dummy"
)

type Response struct {
	Status int         `json:"status,omitempty"`
	Data   dummy.Dummy `json:"data"`
}
