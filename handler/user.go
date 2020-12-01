package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yeo311/echo-boilerplate/model"
	"github.com/yeo311/echo-boilerplate/utils"
)

// CreateUserHandler :
func (h *Handler) CreateUserHandler(c echo.Context) error {
	req := CreateUserRequest{}
	var u model.User
	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}
	if err := h.userStore.CreateUser(&u); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	return c.JSON(http.StatusCreated, map[string]string{"success": "created"})
}
