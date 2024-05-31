package base_repository

import "fmt"

type Filters struct {
	Fields string
	Where  string
	Order  string
	Limit  *int
	Offset *int
}

func (f *Filters) SetWhere(field string, op string, vl any, isString bool) *Filters {
	where := ""
	if f.Where == "" {
		where = "WHERE \n"
	}

	if isString {
		vl = fmt.Sprintf("'%s'", vl)
	}

	f.Where += fmt.Sprintf(" %s %s %s %v \n", where, field, op, vl)
	return f
}

func (f *Filters) And() *Filters {
	f.Where += "AND"
	return f
}

func (f *Filters) Or() *Filters {
	f.Where += "OR"
	return f
}

func (f *Filters) OrderBy(field string, order string) *Filters {
	orderBy := "ORDER BY"
	if f.Order != "" {
		orderBy += ", "
	}

	f.Order += fmt.Sprintf("%s %s %s", orderBy, field, order)
	return f
}
