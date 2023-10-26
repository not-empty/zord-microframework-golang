package dummy

import (
	domain "go-skeleton/application/domain/dummy"
	"go-skeleton/application/services"
)

type Request struct {
	Dummy     domain.Dummy
	Err       error
	validator services.Validator
}

func NewRequest(dummy domain.Dummy, validator services.Validator) Request {
	return Request{
		Dummy:     dummy,
		validator: validator,
	}
}

func (r *Request) Validate() error {
	if err := r.dummyCreateRule(); err != nil {
		return err
	}
	errs := r.validator.ValidateStruct(r.Dummy)
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Request) dummyCreateRule() error {
	return nil
}
