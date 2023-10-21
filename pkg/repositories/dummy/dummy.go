package dummyRepository

import (
	"go-skeleton/application/domain/dummy"
	"go-skeleton/pkg/database"
)

type DummyRepository struct {
	Mysql *database.MySql
}

func NewBaseRepository(mysql *database.MySql) *DummyRepository {
	return &DummyRepository{
		Mysql: mysql,
	}
}

func (repo *DummyRepository) Get(id string) (*dummy.Dummy, error) {
	var Data dummy.Dummy
	err := repo.Mysql.Db.First(&Data, "dummy_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &Data, nil
}

func (repo *DummyRepository) Create(d *dummy.Dummy) error {
	err := repo.Mysql.Db.Create(d).Error
	return err
}

func (repo *DummyRepository) List() (*[]dummy.Dummy, error) {
	var data []dummy.Dummy
	err := repo.Mysql.Db.Limit(100).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *DummyRepository) Edit(d *dummy.Dummy) error {
	err := repo.Mysql.Db.Updates(d).Error
	return err

}

func (repo *DummyRepository) Delete(d *dummy.Dummy) error {
	err := repo.Mysql.Db.Delete(d).Error
	return err
}
