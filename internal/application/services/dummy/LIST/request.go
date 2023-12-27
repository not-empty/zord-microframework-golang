package dummy

type Request struct {
	Data *Data
	Err  error
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
	return nil
}
