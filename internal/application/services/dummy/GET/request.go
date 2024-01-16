package dummy

type Data struct {
	DummyId string `param:"dummy_id"`
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
