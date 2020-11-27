package store

import (
	"github.com/kamva/mgm/v3"
	"github.com/yeo311/echo-boilerplate/model"
)

// UserStore  :
type UserStore struct {
	db *mgm.Collection
}

// NewUserStore :
func NewUserStore(db *mgm.Collection) *UserStore {
	return &UserStore{
		db: db,
	}
}

// CreateUser :
func (us *UserStore) CreateUser(u *model.User) error {
	return us.db.Create(u)
}
