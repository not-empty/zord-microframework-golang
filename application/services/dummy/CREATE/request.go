package dummy

import (
	"errors"
	domain "go-skeleton/application/domain/dummy"
)

type Request struct {
	Dummy domain.Dummy
	Err   error
}

func NewRequest() Request {
	return Request{
		Dummy: domain.Dummy{
			DummyId:   "testando",
			DummyName: "",
		},
	}
}

func (r *Request) Validate() error {
	if err := r.dummyCreateRule(); err != nil {
		return err
	}
	return nil
}

func (r *Request) dummyCreateRule() error {
	if r.Dummy.DummyName != "" {
		return errors.New("invalid_argument")
	}
	return nil
}
