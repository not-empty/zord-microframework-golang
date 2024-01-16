package user

import (
	"go-skeleton/internal/application/domain/user"
	"go-skeleton/internal/application/services"
)

type Service struct {
	services.BaseService
	response   *Response
	repository user.Repository
	idCreator  services.IdCreator
}

func NewService(log services.Logger, repository user.Repository, idCreator services.IdCreator) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
		repository: repository,
		idCreator:  idCreator,
	}
}

func (s *Service) Execute(request Request) {
	if err := request.Validate(); err != nil {
		s.BadRequest(request, err)
		return
	}

	request.Data.UserId = s.idCreator.Create()
	s.produceResponseRule(request.Data)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(data *Data) {
	user := user.User{
		UserId: data.UserId,
	}

	err := s.repository.Create(&user)
	if err != nil {
		s.Error = &services.Error{
			Status:  400,
			Message: "Try again in a few minutes",
			Error:   "Error on request process",
		}
		return
	}
	s.response = &Response{
		Data: user,
	}
}
