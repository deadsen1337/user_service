package user

import (
	"github.com/cockroachdb/pebble"
)

type Repository struct {
	Db *pebble.DB
}

func NewRepository(conn *pebble.DB) *Repository {
	return &Repository{Db: conn}
}

func (r *Repository) Create(id, user []byte) error {
	return r.Db.Set(id, user, nil)
}

func (r *Repository) Update(id, user []byte) error {
	return r.Db.Set(id, user, nil)
}

func (r *Repository) Get(id []byte) ([]byte, error) {
	user, closer, err := r.Db.Get(id)
	if err != nil {
		return nil, err
	}

	if err = closer.Close(); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) Delete(id []byte) error {
	return r.Db.Delete(id, nil)
}

