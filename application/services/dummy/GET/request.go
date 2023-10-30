package dummy

import (
	"errors"
)

type Data struct {
	DummyId string `param:"dummy_id"`
}

type Request struct {
	Data *Data
	Err  error
}

func NewRequest(data *Data) Request {
	return Request{
		Data: data,
	}
}

func (r *Request) Validate() error {
	if err := r.dummyIdRule(); err != nil {
		return err
	}
	return nil
}

func (r *Request) dummyIdRule() error {
	if r.Data.DummyId == `""` {
		return errors.New("invalid_argument")
	}
	return nil
}
