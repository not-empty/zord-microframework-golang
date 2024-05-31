package dummy

import (
	"go-skeleton/internal/application/services"
)

type Data struct {
	DummyId   string
	DummyName string `validate:"required,min=3,max=32"`
	Email     string `validate:"required"`
}

type Request struct {
	Data      *Data
	validator services.Validator
}

func NewRequest(data *Data, validator services.Validator) Request {
	return Request{
		Data:      data,
		validator: validator,
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
