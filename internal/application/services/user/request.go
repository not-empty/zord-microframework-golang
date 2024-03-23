package user

type Data struct {
	UserId string
}

type Request struct {
	Data *Data
	Err  error
}

func NewRequest(data *Data) Request {
	return Request{
		Data: data,
	}
}

func (r *Request) Validate() error {
	if err := r.userCreateRule(); err != nil {
		return err
	}

	return nil
}

func (r *Request) userCreateRule() error {
	// Add validation...
	return nil
}
