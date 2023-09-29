package dummy

import (
	domain "go-skeleton/application/domain/dummy"
)

type Request struct {
	Err   error
	dummy domain.Dummy
}

func NewRequest(dummyId string) Request {
	return Request{
		dummy: domain.Dummy{
			DummyId: dummyId,
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
	return nil
}
