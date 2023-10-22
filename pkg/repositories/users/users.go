package usersRepository

import (
	"go-skeleton/application/domain/users"
	"go-skeleton/pkg/database"
)

type UsersRepository struct {
	Mysql *database.MySql
}

func NewBaseRepository() *UsersRepository {
	return &UsersRepository{}
}

func (repo *UsersRepository) Get(id string) (*users.Users, error) {
	var Data users.Users
	err := repo.Mysql.Db.First(&Data, "users_id = ?", id).Error
	if err != nil {
        return nil, err
    }
	return &Data, nil
}

func (repo *UsersRepository) Create(d *users.Users) error {
	err := repo.Mysql.Db.Create(d).Error
	return err
}

func (repo *UsersRepository) List() (*[]users.Users, error) {
	var data []users.Users
	err := repo.Mysql.Db.Limit(100).Find(&data).Error
	if err != nil {
        return nil, err
    }
	return &data, nil
}

func (repo *UsersRepository) Edit(d *users.Users) error {
	err := repo.Mysql.Db.Updates(d).Error
	return err
}

func (repo *UsersRepository) Delete(d *users.Users) error {
	err := repo.Mysql.Db.Delete(d).Error
	return err
}
