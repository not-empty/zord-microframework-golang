package dummy

import (
	"go-skeleton/internal/application/services"
)

type Data struct {
	ID        string `param:"id"`
	Email     string
	DummyName string
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
