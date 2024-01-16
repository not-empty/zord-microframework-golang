package user

import (
	"go-skeleton/internal/application/providers/pagination"
	"go-skeleton/internal/repositories/base_repository"
)

type User struct {
	UserId   string `gorm:"type:char(26);primarykey" json:"user_id"`
	UserName string `validate:"required,min=3,max=32" json:"user_name"`
}

type Repository interface {
	base_repository.BaseRepository[User]
}

type PaginationProvider interface {
	pagination.IPaginationProvider[User]
}
