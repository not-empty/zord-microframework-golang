package cart

type Data struct {
	CartId string
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
	if err := r.cartCreateRule(); err != nil {
		return err
	}

	return nil
}

func (r *Request) cartCreateRule() error {
	// Add validation...
	return nil
}
