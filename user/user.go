package user

import "github.com/yeo311/echo-boilerplate/model"

// Store : User Store
type Store interface {
	CreateUser(user *model.User) error
}
