package dummy

import (
	"errors"
	domain "go-skeleton/application/domain/dummy"
	"io"
)

type Request struct {
	Dummy domain.Dummy
	Body  io.ReadCloser
	Err   error
}

func NewRequest(Body io.ReadCloser) Request {
	return Request{
		Body: Body,
	}
}

func (r *Request) Validate() error {
	if err := r.dummyCreateRule(); err != nil {
		return err
	}
	return nil
}

func (r *Request) dummyCreateRule() error {

	if r.Dummy.DummyName == "" {
		return errors.New("invalid_argument")
	}
	return nil
}
