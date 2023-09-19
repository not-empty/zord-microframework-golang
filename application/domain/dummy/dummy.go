package dummy

type Dummy struct {
	DummyId   string `gorm:"primarykey"`
	DummyName string `validate:"required,min=3,max=32"`
}

type Repository interface {
	Get(id string) (Dummy, error)
	Create(data *Dummy) bool
	List() []Dummy
	Edit(data *Dummy) bool
	Delete(id *Dummy) bool
}
