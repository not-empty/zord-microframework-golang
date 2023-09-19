package dummy

import (
	domain "go-skeleton/application/domain/dummy"
	"go-skeleton/pkg/validator"
)

type Request struct {
	Dummy     []domain.Dummy
	Err       error
	validator *validator.Validator
}

func NewRequest(validator *validator.Validator) Request {
	return Request{
		validator: validator,
	}
}

func (r *Request) Validate() error {
	errs := r.validator.ValidateStruct(r.Dummy)
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
