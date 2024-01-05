package dummy

import (
	"errors"
	"go-skeleton/internal/application/services"
)

type Data struct {
	Page int `validate:"required"`
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
	errs := r.validator.ValidateStruct(r.Data)

	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	if r.Data.Page < 1 {
		return errors.New("invalid page number")
	}

	return nil
}
