package user

type Data struct {
	UserId string `param:"user_id"`
}

type Request struct {
	Data *Data
	err error
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
	return nil
}
