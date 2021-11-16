package todo

import "github.com/jxlwqq/todo/pkg/dbcontext"

type Repository interface {
	Create(todo *Todo) error
	Update(todo *Todo) error
	Delete(id int64) error
	Get(id int64) (*Todo, error)
	List() ([]*Todo, error)
}

type repository struct {
	db *dbcontext.DB
}

func NewRepository(db *dbcontext.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Create(todo *Todo) error {
	return r.db.Create(todo).Error
}

func (r repository) Update(todo *Todo) error {
	return r.db.Save(todo).Error
}

func (r repository) Delete(id int64) error {
	todo := Todo{ID: id}
	return r.db.Delete(&todo).Error
}

func (r repository) Get(id int64) (*Todo, error) {
	todo := Todo{ID: id}
	err := r.db.First(&todo).Error
	return &todo, err
}

func (r repository) List() ([]*Todo, error) {
	var users []*Todo
	err := r.db.Find(&users).Error
	return users, err
}
