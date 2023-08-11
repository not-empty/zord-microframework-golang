package dummy

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(dummy *Dummy) {

}

func (r *Repository) Get(dummy *Dummy) {

}
