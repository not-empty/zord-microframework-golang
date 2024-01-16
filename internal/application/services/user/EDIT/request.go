package user

import (
	"errors"
	
	"go-skeleton/internal/application/services"
)

type Data struct {
	UserId string `param:"user_id"`
}

type Request struct {
	Data *Data
	
	validator services.Validator
}

func NewRequest(data *Data, validator services.Validator) Request {
	return Request{
		Data: data,
			validator: validator,
	}
}

func (r *Request) Validate() error {
	if err := r.userIdRule(); err != nil {
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

func (r *Request) userIdRule() error {
	if r.Data.UserId == `""` {
		return errors.New("invalid_argument")
	}
	return nil
}
