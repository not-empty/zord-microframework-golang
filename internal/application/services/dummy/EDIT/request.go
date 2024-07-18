package dummy

import (
	"errors"
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/application/services"
)

type Data struct {
	Email     string `validate:"required"`
	DummyName string `validate:"required"`
}

type Request struct {
	ID        string
	Data      *Data
	Domain    *dummy.Dummy
	validator services.Validator
}

func NewRequest(id string, data *Data, validator services.Validator) Request {
	domain := &dummy.Dummy{}
	return Request{
		Data:      data,
		ID:        id,
		validator: validator,
		Domain:    domain,
	}
}

func (r *Request) Validate() error {
	if r.ID == "" {
		return errors.New("invalid id")
	}

	errs := r.validator.ValidateStruct(r.Data)
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
