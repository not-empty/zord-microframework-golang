package dummyRepository

import (
	"go-skeleton/application/domain/dummy"
	"go-skeleton/pkg/database"
)

type DummyRepository struct {
	Mysql *database.MySql
}

func NewBaseRepository() *DummyRepository {
	return &DummyRepository{}
}

func (repo *DummyRepository) Get(id string) (dummy.Dummy, error) {
	var Data dummy.Dummy
	repo.Mysql.Db.First(&Data, "dummy_id = ?", id)
	return Data, nil
}

func (repo *DummyRepository) Create(d *dummy.Dummy) bool {
	repo.Mysql.Db.Create(d)
	return true
}

func (repo *DummyRepository) List() []dummy.Dummy {
	var data []dummy.Dummy
	repo.Mysql.Db.Limit(100).Find(&data)
	return data
}

func (repo *DummyRepository) Edit(d *dummy.Dummy) bool {
	repo.Mysql.Db.Updates(d)
	return true
}

func (repo *DummyRepository) Delete(d *dummy.Dummy) bool {
	repo.Mysql.Db.Delete(d)
	return true
}
