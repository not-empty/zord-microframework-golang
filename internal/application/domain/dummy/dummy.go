package dummy

import (
	"go-skeleton/internal/application/providers/pagination"
	"go-skeleton/internal/repositories/base_repository"
)

type Dummy struct {
	DummyId   string `gorm:"type:char(26);primarykey" json:"dummy_id"`
	DummyName string `validate:"required,min=3,max=32" json:"dummy_name"`
}

type Repository interface {
	base_repository.BaseRepository[Dummy]
}

type PaginationProvider interface {
	pagination.IPaginationProvider[Dummy]
}
