package dummy

import (
	"errors"
	"go-skeleton/internal/application/providers/filters"
)

type Request struct {
	Data    *Data
	Filters filters.Filters
}

type Data struct {
	Page int
	Name string
}

func NewRequest(data *Data, filters filters.Filters) Request {
	return Request{
		Data:    data,
		Filters: filters,
	}
}

func (r *Request) Validate() error {
	if r.Data.Page <= 0 {
		return errors.New("invalid page")
	}

	parseErr := r.Filters.Parse(map[string]filters.FilterData{
		"name": {
			Value:    r.Data.Name,
			IsString: true,
		},
	})

	if parseErr != nil {
		return parseErr
	}

	return nil
}
