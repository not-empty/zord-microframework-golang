package dummy

import "go-skeleton/internal/application/providers/filters"

func (r *Request) SetFiltersRules() error {
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
