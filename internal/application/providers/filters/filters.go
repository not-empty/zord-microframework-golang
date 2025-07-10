package filters

import (
	"errors"
	"fmt"
	"strings"
)

type Filters struct {
	ParsedData []Filter
	FiltersMap map[string]string
}

type Filter struct {
	Field    string
	Operator string
	Value    string
	IsString bool
}

type FilterData struct {
	Value    string
	IsString bool
}

func NewFilters() *Filters {
	return &Filters{
		FiltersMap: map[string]string{
			"eql": "=",
			"neq": "!=",
			"lik": "LIKE",
			"gt":  ">",
			"gte": ">=",
			"lt":  "<",
			"lte": "<=",
		},
	}
}

func (f *Filters) Parse(config map[string]string, data map[string]FilterData) error {
	var filtersData []Filter
	for field, value := range data {
		if value.Value == "" {
			continue
		}
		vl := strings.Split(value.Value, ",")

		if len(vl) != 2 {
			return errors.New("invalid filter definition for: " + field)
		}

		if !f.isValidOperation(config, field, vl[0]) {
			return errors.New("invalid filter operation for: " + field)
		}

		filtersData = append(filtersData, Filter{
			Field:    field,
			Operator: f.FiltersMap[vl[0]],
			Value:    vl[1],
			IsString: value.IsString,
		})
	}

	f.ParsedData = filtersData
	return nil
}

func (f *Filters) isValidOperation(config map[string]string, field string, op string) bool {
	permissions, ok := config[field]
	if !ok {
		return false
	}

	for _, perm := range strings.Split(permissions, ",") {
		if perm == op {
			return true
		}
	}

	return false
}

func (f *Filters) FormatStr(data string) string {
	return fmt.Sprintf("'%s'", data)
}
