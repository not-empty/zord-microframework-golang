package base_repository

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	Fields string
	Where  string
	Order  string
	Limit  *int
	Offset *int
}

func (f *QueryBuilder) SetWhere(field string, op string, vl string, isString bool) *QueryBuilder {
	where := ""
	if f.Where == "" {
		where = "WHERE \n"
	}

	if isString {
		if op == "LIKE" {
			vl = "%" + vl + "%"
		}
		vl = fmt.Sprintf("'%s'", vl)
	}

	f.Where += fmt.Sprintf(" %s %s %s %v \n", where, field, op, vl)
	return f
}

func (f *QueryBuilder) GetWhere() string {
	f.Where = strings.TrimSuffix(f.Where, "AND")
	f.Where = strings.TrimSuffix(f.Where, "OR")
	return f.Where
}

func (f *QueryBuilder) And() *QueryBuilder {
	f.Where += "AND"
	return f
}

func (f *QueryBuilder) Or() *QueryBuilder {
	f.Where += "OR"
	return f
}

func (f *QueryBuilder) OrderBy(field string, order string) *QueryBuilder {
	orderBy := "ORDER BY"
	if f.Order != "" {
		orderBy += ", "
	}

	f.Order += fmt.Sprintf("%s %s %s", orderBy, field, order)
	return f
}
