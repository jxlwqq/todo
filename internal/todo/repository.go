package todo

import "github.com/jxlwqq/todo/pkg/dbcontext"

type Repository interface {
	Create(item *Item) error
	Update(item *Item) error
	Delete(id int64) error
	Get(id int64) (*Item, error)
	List() ([]*Item, error)
}

type repository struct {
	db *dbcontext.DB
}

func NewRepository(db *dbcontext.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Create(item *Item) error {
	return r.db.Create(item).Error
}

func (r repository) Update(item *Item) error {
	return r.db.Save(item).Error
}

func (r repository) Delete(id int64) error {
	item := Item{ID: id}
	return r.db.Delete(&item).Error
}

func (r repository) Get(id int64) (*Item, error) {
	item := Item{ID: id}
	err := r.db.First(&item).Error
	return &item, err
}

func (r repository) List() ([]*Item, error) {
	var users []*Item
	err := r.db.Find(&users).Error
	return users, err
}
