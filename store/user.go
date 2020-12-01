package store

import (
	"context"
	"log"

	"github.com/kamva/mgm/v3"
	"github.com/yeo311/echo-boilerplate/model"
	"go.mongodb.org/mongo-driver/bson"
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

// GetByEmail :
func (us *UserStore) GetByEmail(email string) (*model.User, error) {
	var u *model.User
	if err := us.db.FindOne(context.TODO(), bson.M{"email": email}).Decode(u); err != nil {
		log.Println(err)
		return nil, err
	}
	return u, nil
}
