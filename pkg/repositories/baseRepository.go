package repositories

type BaseRepository[T any] struct {
}

func NewBaseRepository[T any]() *BaseRepository[T] {
	return &BaseRepository[T]{}
}

func (repo *BaseRepository[T]) Get(id string) T {
	var teste T
	return teste
}

func (repo *BaseRepository[T]) Create(*T) bool {
	return true
}

func (repo *BaseRepository[T]) List() []T {
	var teste []T
	return teste
}

func (repo *BaseRepository[T]) Edit(*T) bool {
	return true
}
