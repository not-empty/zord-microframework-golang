package dummy

import (
	"go-skeleton/internal/application/services"
)

type Data struct {
	DummyName string `validate:"required,min=3,max=32"`
	Email     string `validate:"required"`
}

type Request struct {
	Data      *Data
	Client    string
	validator services.Validator
}

func NewRequest(data *Data, validator services.Validator, client string) Request {
	return Request{
		Client:    client,
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
