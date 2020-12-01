package utils

import "github.com/labstack/echo/v4"

// Error :
type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

// NewError :
func NewError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Errors["body"] = v.Message
	default:
		e.Errors["body"] = v.Error()
	}
	return e
}
