package user

import (
	"go-skeleton/internal/application/domain/user"
	"go-skeleton/internal/application/services"
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
	s.Logger.Debug("Hello Im User Server!")
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
	userData, err := s.repository.Get(data.UserId, "user_id")
	if err != nil {
		s.Error = &services.Error{
			Status:  400,
			Message: "Try again in a few minutes",
			Error:   "Error on request process",
		}
		return
	}
	s.response = &Response{
		Data: userData,
	}
}
