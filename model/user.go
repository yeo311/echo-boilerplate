package model

import (
	"github.com/kamva/mgm/v3"
)

// User :
type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `bson:"email"`
	Password         string `bson:"password"`
	Name             string `bson:"name"`
}

// NewUser :
func NewUser(email, password, name string) *User {
	return &User{
		Email:    email,
		Password: password,
		Name:     name,
	}
}
