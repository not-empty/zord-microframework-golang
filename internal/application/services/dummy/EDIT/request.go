package dummy

import (
	"errors"
	"go-skeleton/internal/application/services"
)

type Data struct {
	DummyId   string `param:"dummy_id"`
	DummyName string `json:"dummy_name"`
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
	if err := r.dummyIdRule(); err != nil {
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

func (r *Request) dummyIdRule() error {
	if r.Data.DummyId == `""` {
		return errors.New("invalid_argument")
	}
	return nil
}
