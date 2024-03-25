package base_repository

import (
	"errors"
	"fmt"
	"github.com/fatih/structs"
	"github.com/jmoiron/sqlx"
	"strings"
)

type BaseRepository[dom Domain] interface {
	Get(id string, field string) (*dom, error)
	Create(data dom) error
	List(limit int, offset int) (*[]dom, error)
	Edit(data dom, field string, value string) error
	Delete(data dom) error
	Count() (int64, error)
}

type Domain interface {
	Schema() string
}

type BaseRepo[dom Domain] struct {
	Mysql  *sqlx.DB
	fields []string
}

func NewBaseRepository[dom Domain](mysql *sqlx.DB) *BaseRepo[dom] {
	do := new(dom)
	listFields := []string{}
	fields := structs.Fields(&do)
	for _, field := range fields {
		tag := field.Tag("db")
		listFields = append(listFields, tag)
	}

	return &BaseRepo[dom]{
		Mysql:  mysql,
		fields: listFields,
	}
}

func (repo *BaseRepo[Domain]) Get(value string, field string) (*Domain, error) {
	var Data Domain
	row := repo.Mysql.QueryRowx(
		fmt.Sprintf(
			"SELECT %s FROM %s WHERE %s = ?",
			strings.Join(repo.fields, ", "),
			Data.Schema(),
			field,
		),
		value,
	)
	err := row.StructScan(&Data)
	if err != nil {
		return nil, err
	}
	return &Data, nil
}

func (repo *BaseRepo[Row]) Create(d Row) error {
	tx, err := repo.Mysql.Beginx()
	if err != nil {
		return err
	}

	namedValues := []string{}
	for _, field := range repo.fields {
		namedValues = append(namedValues, ":"+field)
	}
	res, err := tx.NamedExec(
		fmt.Sprintf(
			`INSERT INTO %s (%s) VALUES (%s)`,
			d.Schema(),
			strings.Join(repo.fields, ", "),
			strings.Join(namedValues, ", "),
		),
		d,
	)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if affected < 1 {
		return errors.New("erro ao inserir registro")
	}
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repo *BaseRepo[Row]) List(limit int, offset int) (*[]Row, error) {
	var data []Row
	var value Row
	rows, err := repo.Mysql.Queryx(
		fmt.Sprintf(
			`SELECT %s FROM %s LIMIT %v OFFSET %v`,
			strings.Join(repo.fields, ", "),
			value.Schema(),
			limit,
			offset,
		),
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.StructScan(&value)
		if err != nil {
			return nil, err
		}
		data = append(data, value)
	}
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *BaseRepo[Row]) Edit(d Row, field string, value string) error {
	tx, err := repo.Mysql.Beginx()
	if err != nil {
		return err
	}

	namedValues := []string{}
	for _, field := range repo.fields {
		namedValues = append(namedValues, field+" = :"+field)
	}
	res, err := tx.NamedExec(
		fmt.Sprintf(
			`UPDATE %s SET %s WHERE %s = %s`,
			d.Schema(),
			strings.Join(namedValues, ", "),
			field,
			value,
		),
		d,
	)

	fmt.Println(fmt.Sprintf(
		`UPDATE %s SET %s WHERE %s = '%s'`,
		d.Schema(),
		strings.Join(namedValues, ", "),
		field,
		value,
	))
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if affected < 1 {
		return errors.New("erro ao editar registro")
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repo *BaseRepo[Row]) Delete(d Row) error {
	//err := repo.Mysql.Db.Delete(d).Error
	return nil
}

func (repo *BaseRepo[Row]) Count() (int64, error) {
	//var count int64
	//err := repo.Mysql.Db.Model(new(Row)).Count(&count).Error
	//if err != nil {
	//	return 0, err
	//}
	return 0, nil
}
