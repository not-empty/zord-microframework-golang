package dummy

import (
	"errors"
	"go-skeleton/application/services"
)

type RequestDTO struct {
	DummyId   string `param:"dummy_id"`
	DummyName string `json:"dummy_name"`
}

type Request struct {
	DTO       *RequestDTO
	validator services.Validator
}

func NewRequest(dto *RequestDTO, validator services.Validator) Request {
	return Request{
		DTO:       dto,
		validator: validator,
	}
}

func (r *Request) Validate() error {
	if err := r.dummyIdRule(); err != nil {
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

func (r *Request) dummyIdRule() error {
	if r.DTO.DummyId == `""` {
		return errors.New("invalid_argument")
	}
	return nil
}
