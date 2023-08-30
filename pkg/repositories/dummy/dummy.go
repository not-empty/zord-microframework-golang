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
	repo.Mysql.Db.Find(&Data, id)
	return Data, nil
}

func (repo *DummyRepository) Create(data *dummy.Dummy) bool {
	repo.Mysql.Db.Create(data)
	return true
}

func (repo *DummyRepository) List() []dummy.Dummy {
	var data []dummy.Dummy
	repo.Mysql.Db.Find(&data).Limit(100)
	return data
}

func (repo *DummyRepository) Edit(data *dummy.Dummy) bool {
	repo.Mysql.Db.Updates(data)
	return true
}

func (repo *DummyRepository) Delete(data *dummy.Dummy) bool {
	repo.Mysql.Db.Delete(data)
	return true
}
