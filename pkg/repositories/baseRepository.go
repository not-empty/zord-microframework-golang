package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type BaseRepository[T any] struct {
	db     sql.DB
	domain string
}

func NewBaseRepository[T any]() *BaseRepository[T] {
	return &BaseRepository[T]{}
}

func (repo *BaseRepository[T]) Get(id string) (*T, error) {
	var domain T
	result, err := repo.db.Query("SELECT * FROM " + repo.domain + " LIMIT 1")
	if err != nil {
		return nil, err
	}
	err = result.Scan(&domain)
	if err != nil {
		return nil, err
	}
	return &domain, nil
}

func (repo *BaseRepository[T]) Create(data *T) bool {
	// tx, err := repo.db.Begin()
	// if err != nil {
	// 	return false
	// }

	return true
}

func (repo *BaseRepository[T]) List() []T {
	var teste []T
	return teste
}

func (repo *BaseRepository[T]) Edit(*T) bool {
	return true
}

func (repo *BaseRepository[T]) Delete(id string) bool {
	return true
}
