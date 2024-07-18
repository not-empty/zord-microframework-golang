package dummy

import (
	"errors"
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/application/providers/filters"
)

type Request struct {
	Data    *Data
	Filters *filters.Filters
	Domain  *dummy.Dummy
}

type Data struct {
	Page  int
	Name  string
	Email string
}

func NewRequest(data *Data, filters *filters.Filters) Request {
	domain := &dummy.Dummy{}
	return Request{
		Data:    data,
		Filters: filters,
		Domain:  domain,
	}
}

func (r *Request) Validate() error {
	if r.Data.Page <= 0 {
		return errors.New("invalid page")
	}

	return nil
}
