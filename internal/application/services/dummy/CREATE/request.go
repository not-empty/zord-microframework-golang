package dummy

import (
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/application/services"
)

type Data struct {
	DummyName string `validate:"required,min=3,max=32"`
	Email     string `validate:"required"`
}

type Request struct {
	Data      *Data
	Domain    *dummy.Dummy
	validator services.Validator
}

func NewRequest(data *Data, validator services.Validator) *Request {
	domain := &dummy.Dummy{}
	return &Request{
		Data:      data,
		validator: validator,
		Domain:    domain,
	}
}

func (r *Request) Validate() error {
	errs := r.validator.ValidateStruct(r.Data)
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
