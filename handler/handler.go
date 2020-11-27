package handler

import (
	"github.com/yeo311/echo-boilerplate/user"
)

// Handler :
type Handler struct {
	userStore user.Store
}

// NewHandler :
func NewHandler(us user.Store) *Handler {
	return &Handler{
		userStore: us,
	}
}
