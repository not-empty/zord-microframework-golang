package dummy

import (
	"errors"
	domain "go-skeleton/application/domain/dummy"
	"io"
)

type Request struct {
	Dummy   domain.Dummy
	Body    io.ReadCloser
	Dummyid string
}

func NewRequest(body io.ReadCloser, dummyId string) Request {
	return Request{
		Body:    body,
		Dummyid: dummyId,
	}
}

func (r *Request) Validate() error {
	if err := r.dummyIdRule(); err != nil {
		return err
	}
	return nil
}

func (r *Request) dummyIdRule() error {
	if r.Dummy.DummyId == `""` {
		return errors.New("invalid_argument")
	}
	return nil
}
