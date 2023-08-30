package dummy

import (
	"errors"
	domain "go-skeleton/application/domain/dummy"
)

type Request struct {
	Dummy   domain.Dummy
	Dummyid string
}

func NewRequest(dummy domain.Dummy, dummyId string) Request {
	req := Request{
		Dummy: dummy,
	}
	req.Dummy.DummyId = dummyId
	return req
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
