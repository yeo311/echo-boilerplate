package handler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/yeo311/echo-boilerplate/model"
)

// CreateUserRequest :
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (r *CreateUserRequest) bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	u.Email = r.Email
	u.Password = r.Password
	u.Name = r.Name

	h, err := u.HashPassword(u.Password)
	if err != nil {
		log.Println(err)
		return err
	}
	u.Password = h

	return nil
}
