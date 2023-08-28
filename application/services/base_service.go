package services

import (
	"net/http"
)

type Logger interface {
	Debug(Message string, Context ...string)
	Info(Message string, Context ...string)
	Warning(Message string, Context ...string)
	Error(Error error, Context ...string)
	Critical(Error error, Context ...string)
}

type Repository[T any] interface {
	Get(string) (*T, error)
	Create(*T) bool
	List() []T
	Edit(*T) bool
	Delete(string) bool
}

type Error struct {
	Status  int    `json:"-"`
	Message string `json:"-"`
	Error   string `json:"error"`
}

type BaseService struct {
	Logger Logger
	Error  *Error
}

type Request interface {
}

func (bs *BaseService) BadRequest(request Request, err error) {
	bs.Error = &Error{
		Status:  http.StatusBadRequest,
		Message: http.StatusText(http.StatusBadRequest),
		Error:   err.Error(),
	}
}
