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

type ICrypt interface {
	GenerateHash(target string) (string, error)
	CompareHash(hashed string, target string) bool
}

type Signer interface {
	Sign(signingString string) (string, error)
	Verify(jwtStr string) (bool, error)
}

type IdCreator interface {
	Create() string
}

type Validator interface {
	ValidateStruct(modelData any) []error
}

type Error struct {
	Status  int
	Message interface{}
	Error   string
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

func (bs *BaseService) InternalServerError(msg string, err error) {
	bs.Logger.Error(err)
	bs.errorHandler(http.StatusInternalServerError, msg)
}

func (bs *BaseService) NotFound(msg string) {
	bs.errorHandler(http.StatusNotFound, msg)
}

func (bs *BaseService) BadRequest(msg string) {
	bs.errorHandler(http.StatusBadRequest, msg)
}

func (bs *BaseService) UnprocessableEntity(msg string) {
	bs.errorHandler(http.StatusUnprocessableEntity, msg)
}
