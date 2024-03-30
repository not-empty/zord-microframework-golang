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

type IdCreator interface {
	Create() string
}

type Validator interface {
	ValidateStruct(modelData any) []error
}

type Error struct {
	Status  int         `json:"-"`
	Message interface{} `json:"message"`
	Error   string      `json:"error"`
}

type BaseService struct {
	Logger Logger
	Error  *Error
	Ulid   IdCreator
}

type Request interface {
}

func (bs *BaseService) errorHandler(httpStatus int, errMsg interface{}) {
	bs.Error = &Error{
		Status:  httpStatus,
		Message: errMsg,
		Error:   http.StatusText(httpStatus),
	}
}

func (bs *BaseService) CustomError(status int, err interface{}) {
	bs.errorHandler(status, err)
}

func (bs *BaseService) InternalServerError(errMsg string) {
	bs.errorHandler(http.StatusInternalServerError, errMsg)
}

func (bs *BaseService) NotFound(errMsg string) {
	bs.errorHandler(http.StatusNotFound, errMsg)
}

func (bs *BaseService) BadRequest(errMsg string) {
	bs.errorHandler(http.StatusBadRequest, errMsg)
}

func (bs *BaseService) UnprocessableEntity(errMsg string) {
	bs.errorHandler(http.StatusUnprocessableEntity, errMsg)
}
