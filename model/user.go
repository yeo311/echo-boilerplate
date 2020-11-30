package model

import (
	"errors"

	"github.com/kamva/mgm/v3"
	"golang.org/x/crypto/bcrypt"
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

// HashPassword :
func (u *User) HashPassword(plain string) (string, error) {
	if plain == "" {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}
