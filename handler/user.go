package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yeo311/echo-boilerplate/model"
)

// CreateUserHandler :
func (h *Handler) CreateUserHandler(c echo.Context) error {
	req := CreateUserRequest{}
	var u model.User
	if err := req.bind(c, &u); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.userStore.CreateUser(&u); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"success": "created"})
}
