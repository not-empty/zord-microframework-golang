package dummy

type Dummy struct {
	DummyId   string `gorm:"type:char(26);primarykey" json:"dummy_id"`
	DummyName string `validate:"required,min=3,max=32" json:"dummy_name"`
}

type Repository interface {
	Get(id string) (*Dummy, error)
	Create(d *Dummy) error
	List() (*[]Dummy, error)
	Edit(d *Dummy) error
	Delete(d *Dummy) error
}
