package dummy

type Data struct {
	ID string `param:"id"`
}

type Request struct {
	Data *Data
}

func NewRequest(data *Data) Request {
	return Request{
		Data: data,
	}
}

func (r *Request) Validate() error {
	return nil
}
