package dummy

import (
	"go-skeleton/application/services"
)

type RequestDTO struct {
	DummyId   string
	DummyName string `validate:"required,min=3,max=32" json:"dummy_name"`
}

type Request struct {
	DTO       *RequestDTO
	Err       error
	validator services.Validator
}

func NewRequest(dto *RequestDTO, validator services.Validator) Request {
	return Request{
		DTO:       dto,
		validator: validator,
	}
}

func (r *Request) Validate() error {
	if err := r.dummyCreateRule(); err != nil {
		return err
	}
	errs := r.validator.ValidateStruct(r.DTO)
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
