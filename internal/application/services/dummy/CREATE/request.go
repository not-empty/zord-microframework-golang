package dummy

import (
	"go-skeleton/internal/application/services"
)

type Data struct {
	DummyId   string
	DummyName string `validate:"required,min=3,max=32" json:"dummy_name"`
}

type Request struct {
	Data      *Data
	Err       error
	validator services.Validator
}

func NewRequest(data *Data, validator services.Validator) Request {
	return Request{
		Data:      data,
		validator: validator,
	}
}

func (r *Request) Validate() error {
	if err := r.dummyCreateRule(); err != nil {
		return err
	}
	errs := r.validator.ValidateStruct(r.Data)
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
