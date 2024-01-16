package user

import (
	"errors"
)

type Data struct {
	UserId string `param:"user_id"`
}


type Request struct {
	Data *Data
	Err error
}

func NewRequest(data *Data) Request {
	return Request{
		Data: data,
	}
}

func (r *Request) Validate() error {
	if err := r.userIdRule(); err != nil {
		return err
	}
	return nil
}

func (r *Request) userIdRule() error {
	if r.Data.UserId == `""` {
		return errors.New("invalid_argument")
	}
	return nil
}
