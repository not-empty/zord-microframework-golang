package dummy

type Request struct {
	Err error
}

func NewRequest() Request {
	return Request{}
}

func (r *Request) Validate() error {
	return nil
}
