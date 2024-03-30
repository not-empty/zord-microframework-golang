package dummy

import (
	"errors"
)

type Request struct {
	Data *Data
}

type Data struct {
	Page int
}

func NewRequest(data *Data) Request {
	return Request{
		Data: data,
	}
}

func (r *Request) Validate() error {
	if r.Data.Page <= 0 {
		return errors.New("invalid page")
	}
	return nil
}
