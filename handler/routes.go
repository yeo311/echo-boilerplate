package handler

import "github.com/labstack/echo/v4"

// Register :
func (h *Handler) Register(v1 *echo.Group) {
	user := v1.Group("/user")
	user.POST("", h.CreateUserHandler)
}
