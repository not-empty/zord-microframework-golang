package base_repository

import (
	"errors"
	"fmt"
	"go-skeleton/internal/application/providers/filters"
	"strings"

	"github.com/fatih/structs"
	"github.com/jmoiron/sqlx"
)

type BaseRepository[dom Domain] interface {
	Get(domain dom, field string, value string) (*dom, error)
	Create(data dom, tx *sqlx.Tx, autoCommit bool) error
	List(domain dom, limit int, offset int) (*[]dom, error)
	Search(domain dom, field string, value string) (*[]dom, error)
	Edit(data dom, field string, value string) (int, error)
	Delete(domain dom, field string, values string) error
	Count(Data dom) (int64, error)
	InitTX() (*sqlx.Tx, error)
	Commit(tx *sqlx.Tx) error
	Rollback(tx *sqlx.Tx, err error) error
	NewFilters() *QueryBuilder
}

type Domain interface {
	Schema() string
	GetFilters() filters.Filters
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
		if !field.IsExported() {
			continue
		}
		tag := field.Tag("db")
		listFields = append(listFields, tag)
	}

	return &BaseRepo[dom]{
		Mysql:  mysql,
		fields: listFields,
	}
}

func (repo *BaseRepo[Domain]) InitTX() (*sqlx.Tx, error) {
	tx, err := repo.Mysql.Beginx()
	return tx, err
}

func (repo *BaseRepo[Domain]) Commit(tx *sqlx.Tx) error {
	return tx.Commit()
}

func (repo *BaseRepo[Domain]) Rollback(tx *sqlx.Tx, err error) error {
	rollbErr := tx.Rollback()
	if rollbErr != nil {
		return rollbErr
	}
	return err
}

func (repo *BaseRepo[Domain]) NewFilters() *QueryBuilder {
	return &QueryBuilder{
		Fields: "",
		Where:  "",
		Order:  "",
		Limit:  nil,
		Offset: nil,
	}
}

func (repo *BaseRepo[Domain]) Get(Data Domain, field string, value string) (*Domain, error) {
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

func (repo *BaseRepo[Domain]) Create(Data Domain, tx *sqlx.Tx, autoCommit bool) error {
	namedValues := []string{}
	for _, field := range repo.fields {
		namedValues = append(namedValues, ":"+field)
	}
	res, err := tx.NamedExec(
		fmt.Sprintf(
			`INSERT INTO %s (%s) VALUES (%s)`,
			Data.Schema(),
			strings.Join(repo.fields, ", "),
			strings.Join(namedValues, ", "),
		),
		Data,
	)

	if err != nil {
		return repo.Rollback(tx, err)
	}

	affected, err := res.RowsAffected()
	if affected < 1 {
		return repo.Rollback(tx, errors.New("error on create, rows not affected"))
	}

	if err != nil {
		return repo.Rollback(tx, err)
	}

	if autoCommit {
		err = tx.Commit()
		if err != nil {
			return repo.Rollback(tx, err)
		}
	}

	return nil
}

func (repo *BaseRepo[Domain]) List(Data Domain, limit int, offset int) (*[]Domain, error) {
	var results []Domain
	queryBuilder := repo.NewFilters()
	f := Data.GetFilters()
	for _, data := range f.ParsedData {
		queryBuilder.SetWhere(data.Field, data.Operator, data.Value, data.IsString)
		queryBuilder.And()
	}

	where := queryBuilder.GetWhere()

	rows, err := repo.Mysql.Queryx(
		fmt.Sprintf(
			`SELECT
						%s
					FROM
						%s
					%s
					LIMIT %v OFFSET %v`,
			strings.Join(repo.fields, ", "),
			Data.Schema(),
			where,
			limit,
			offset,
		),
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.StructScan(&Data)
		if err != nil {
			return nil, err
		}
		results = append(results, Data)
	}
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (repo *BaseRepo[Domain]) Edit(Data Domain, field string, value string) (int, error) {
	tx, err := repo.Mysql.Beginx()
	if err != nil {
		return 0, err
	}

	namedValues := []string{}
	for _, field := range repo.fields {
		namedValues = append(namedValues, "`"+field+"`"+" = :"+field)
	}

	query, args, bindErr := tx.BindNamed(
		fmt.Sprintf(
			`UPDATE %s SET %s WHERE %s = '%s'`,
			Data.Schema(),
			strings.Join(namedValues, ", "),
			field,
			value,
		),
		&Data,
	)

	if bindErr != nil {
		return 0, repo.Rollback(tx, bindErr)
	}

	res, execErr := repo.Mysql.Exec(query, args...)
	if execErr != nil {
		return 0, repo.Rollback(tx, execErr)
	}

	affected, rowsAffErr := res.RowsAffected()
	if rowsAffErr != nil {
		return int(affected), repo.Rollback(tx, rowsAffErr)
	}

	if affected < 1 {
		repo.Rollback(tx, nil)
		return int(affected), nil
	}

	commitErr := repo.Commit(tx)
	if commitErr != nil {
		return int(affected), repo.Rollback(tx, commitErr)
	}

	return int(affected), nil
}

func (repo *BaseRepo[Domain]) Delete(Data Domain, field string, value string) error {
	exec, err := repo.Mysql.Exec(
		fmt.Sprintf(
			`DELETE FROM %s WHERE %s = '%s'`,
			Data.Schema(),
			field,
			value,
		),
	)
	if err != nil {
		return err
	}
	rowsAffected, err := exec.RowsAffected()
	if rowsAffected < 1 {
		return errors.New("nothing deleted")
	}
	return err
}

func (repo *BaseRepo[Domain]) Count(Data Domain) (int64, error) {
	var count int64
	queryBuilder := repo.NewFilters()
	f := Data.GetFilters()
	for _, data := range f.ParsedData {
		queryBuilder.SetWhere(data.Field, data.Operator, data.Value, data.IsString)
		queryBuilder.And()
	}

	where := queryBuilder.GetWhere()
	err := repo.Mysql.Get(&count, "SELECT count(1) FROM "+Data.Schema()+" "+where)
	return count, err
}

func (repo *BaseRepo[Domain]) Search(Data Domain, field string, value string) (*[]Domain, error) {
	var data []Domain
	queryErr := repo.Mysql.Select(
		&data,
		fmt.Sprintf(
			`SELECT %s FROM %s WHERE %s like ? LIMIT 25`,
			strings.Join(repo.fields, ", "),
			Data.Schema(),
			field,
		),
		"%"+value+"%",
	)
	if queryErr != nil {
		return nil, queryErr
	}

	return &data, nil
}
