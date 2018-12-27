package errors

import (
	"fmt"
	"net/http"
)

// simple consumer facing error type

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	StatusCode  int    `json:"-"`
}

const (
	ErrInternalServer = "SERVER_ERROR"
	ErrBadRequest     = "BAD_REQUEST"
	ErrUnauthorized   = "UNAUTHORIZED"
)

func (p *Error) Error() string {
	if p == nil {
		return ""
	}
	if p.Description == "" {
		return p.Code
	}
	if p.Code == "" {
		return p.Description
	}
	return fmt.Sprintf("%s: %s", p.Code, p.Description)
}

func New(code string, message string, statusCode int) *Error {
	return errorFactory(code, message, statusCode)
}

func ServerError() *Error {
	return errorFactory(ErrInternalServer, "Sorry, something went wrong.", http.StatusInternalServerError)
}

func BadRequest(message string) *Error {
	return errorFactory(ErrBadRequest, message, http.StatusBadRequest)
}

func Unauthorized(message string) *Error {
	return errorFactory(ErrUnauthorized, message, http.StatusUnauthorized)
}

func errorFactory(code string, message string, statusCode int) *Error {
	err := &Error{
		Code:        ErrInternalServer,
		Description: message,
		StatusCode:  statusCode,
	}
	if len(code) > 0 {
		err.Code = code
	}
	return err
}
