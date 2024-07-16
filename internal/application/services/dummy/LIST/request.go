package dummy

import (
	"errors"
	"go-skeleton/internal/application/providers/filters"
)

type Request struct {
	Data    *Data
	Filters filters.Filters
	Client  string
}

type Data struct {
	Page  int
	Name  string
	Email string
}

func NewRequest(data *Data, filters filters.Filters, client string) Request {
	return Request{
		Data:    data,
		Filters: filters,
		Client:  client,
	}
}

func (r *Request) Validate() error {
	if r.Data.Page <= 0 {
		return errors.New("invalid page")
	}

	parseErr := r.Filters.Parse(
		map[string]string{
			"name":  "eql,lik",
			"email": "lik",
		},
		map[string]filters.FilterData{
			"name": {
				Value:    r.Data.Name,
				IsString: true,
			},
			"email": {
				Value:    r.Data.Email,
				IsString: true,
			},
		})

	if parseErr != nil {
		return parseErr
	}

	return nil
}
