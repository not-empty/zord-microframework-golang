package UserRepository

import (
	"go-skeleton/internal/application/domain/user"
	"go-skeleton/internal/repositories/base_repository"
	"go-skeleton/pkg/database"
)

type UserRepository struct {
	base_repository.BaseRepository[user.User]
}

func NewUserRepo(mysql *database.MySql) *UserRepository {
	return &UserRepository{
		BaseRepository: base_repository.NewBaseRepository[user.User](mysql),
	}
}
