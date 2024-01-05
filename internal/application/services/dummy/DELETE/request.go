package dummy

import "go-skeleton/internal/application/services"

type Data struct {
	DummyId string `param:"dummy_id" validate:"required,ulid"`
}

type Request struct {
	Data      *Data
	Err       error
	validator services.Validator
}

func NewRequest(data *Data, validator services.Validator) Request {
	return Request{
		Data: data,
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
	return nil
}
