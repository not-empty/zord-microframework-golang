package dummy

import (
	"errors"
	domain "go-skeleton/application/domain/dummy"
	"go-skeleton/pkg/validator"
)

type Request struct {
	Dummy     domain.Dummy
	Dummyid   string
	validator *validator.Validator
}

func NewRequest(dummy domain.Dummy, dummyId string, validator *validator.Validator) Request {
	req := Request{
		Dummy:     dummy,
		validator: validator,
	}
	req.Dummy.DummyId = dummyId
	return req
}

func (r *Request) Validate() error {
	if err := r.dummyIdRule(); err != nil {
		return err
	}
	errs := r.validator.ValidateStruct(r.Dummy)
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Request) dummyIdRule() error {
	if r.Dummy.DummyId == `""` {
		return errors.New("invalid_argument")
	}
	return nil
}
