package user

import (
	"go-skeleton/internal/application/domain/user"
	"go-skeleton/internal/application/services"

	"net/http"
)

type Service struct {
	services.BaseService
	response   *Response
	repository user.Repository
}

func NewService(log services.Logger, repository user.Repository) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
		repository: repository,
	}
}

func (s *Service) Execute(request Request) {
	if err := request.Validate(); err != nil {
		s.BadRequest(request, err)
		return
	}

	s.produceResponseRule(request.Data)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(data *Data) {
	userData := user.User{
		UserId:   data.UserId,
	}

	affectedRows, err := s.repository.Edit(&userData)
	if err != nil {
		s.Error = &services.Error{
			Status:  http.StatusInternalServerError,
			Message: "Try again in a few minutes",
			Error:   "Error on request process",
		}
		return
	}

	if affectedRows < 1 {
		s.Error = &services.Error{
			Status:  http.StatusNotFound,
			Message: "Try again in a few minutes",
			Error:   "Error on request process",
		}
		return
	}

	s.response = &Response{
		Data: data,
	}
}
