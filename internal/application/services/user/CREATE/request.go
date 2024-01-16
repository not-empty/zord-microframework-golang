package user

import (
	
	"go-skeleton/internal/application/services"
)

type Data struct {
	UserId string
}

type Request struct {
	Data *Data
	Err error
	
	validator services.Validator
}

func NewRequest(data *Data, validator services.Validator) Request {
	return Request{
		Data: data,
			validator: validator,
	}
}

func (r *Request) Validate() error {
	if err := r.userCreateRule(); err != nil {
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

func (r *Request) userCreateRule() error {
	// Add validation...
	return nil
}
