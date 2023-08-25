package dummy

import (
	"errors"
	domain "go-skeleton/application/domain/dummy"
)

type Request struct {
	Dummy domain.Dummy
	// err   error
}

func NewRequest(dummy string) Request {
	return Request{
		Dummy: domain.Dummy{
			DummyId: dummy,
		},
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
