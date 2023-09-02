package dummy

type Dummy struct {
	DummyId   string `gorm:"primarykey"`
	DummyName string
}

type Repository interface {
	Get(id string) (Dummy, error)
	Create(data *Dummy) bool
	List() []Dummy
	Edit(data *Dummy) bool
	Delete(id *Dummy) bool
}
