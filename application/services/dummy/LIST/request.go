package dummy

import (
	domain "go-skeleton/application/domain/dummy"
)

type Request struct {
	Dummy []domain.Dummy
	// err   error
}

func NewRequest() Request {
	return Request{}
}

func (r *Request) Validate() error {
	return nil
}
