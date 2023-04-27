package user

import (
	"encoding/json"
	"github.com/cockroachdb/pebble"
	"user_service/internal/model"
)

type Repository struct {
	Db *pebble.DB
}

func NewRepository(conn *pebble.DB) *Repository {
	return &Repository{Db: conn}
}

func (r *Repository) Create(id []byte, user model.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return r.Db.Set(id, data, nil)
}

func (r *Repository) Update(id []byte, user model.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return r.Db.Set(id, data, nil)
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

