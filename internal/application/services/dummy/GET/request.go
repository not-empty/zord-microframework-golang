package dummy

type Data struct {
	ID string `param:"id"`
}

type Request struct {
	Data   *Data
	Client string
}

func NewRequest(data *Data, client string) Request {
	return Request{
		Data:   data,
		Client: client,
	}
}

func (r *Request) Validate() error {
	return nil
}
