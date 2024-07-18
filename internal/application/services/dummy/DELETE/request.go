package dummy

import "go-skeleton/internal/application/domain/dummy"

type Data struct {
	ID string `param:"id"`
}

type Request struct {
	Data   *Data
	Domain *dummy.Dummy
}

func NewRequest(data *Data) Request {
	domain := &dummy.Dummy{}
	return Request{
		Data:   data,
		Domain: domain,
	}
}

func (r *Request) Validate() error {
	return nil
}
