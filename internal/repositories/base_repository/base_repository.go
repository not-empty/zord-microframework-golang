package base_repository

import (
	"go-skeleton/pkg/database"
)

type BaseRepository[Row any] interface {
	Get(id string, field string) (*Row, error)
	Create(*Row) error
	List(limit int, offset int) (*[]Row, error)
	Edit(*Row) error
	Delete(*Row) error
	Count() (int64, error)
}

type BaseRepo[Row any] struct {
	Mysql *database.MySql
}

func NewBaseRepository[Row any](mysql *database.MySql) *BaseRepo[Row] {
	return &BaseRepo[Row]{
		Mysql: mysql,
	}
}

func (repo *BaseRepo[Row]) Get(id string, field string) (*Row, error) {
	var Data Row
	err := repo.Mysql.Db.First(&Data, field+" = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &Data, nil
}

func (repo *BaseRepo[Row]) Create(d *Row) error {
	err := repo.Mysql.Db.Create(d).Error
	return err
}

func (repo *BaseRepo[Row]) List(limit int, offset int) (*[]Row, error) {
	var data []Row
	err := repo.Mysql.Db.Limit(limit).Offset(offset).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *BaseRepo[Row]) Edit(d *Row) error {
	err := repo.Mysql.Db.Updates(d).Error
	return err
}

func (repo *BaseRepo[Row]) Delete(d *Row) error {
	err := repo.Mysql.Db.Delete(d).Error
	return err
}

func (repo *BaseRepo[Row]) Count() (int64, error) {
	var count int64
	err := repo.Mysql.Db.Model(new(Row)).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
