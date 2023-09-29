package dummy

import (
	"go-skeleton/application/domain/dummy"
	"go-skeleton/application/services"
	"net/http"
)

type Service struct {
	services.BaseService
	response   *Response
	repository dummy.Repository
}

func NewService(log services.Logger, repository dummy.Repository) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
		repository: repository,
	}
}

func (s *Service) Execute(request Request) {
	s.Logger.Debug("Hello Im Dummy Server!")
	if err := request.Validate(); err != nil {
		s.BadRequest(request, err)
		return
	}

	s.produceResponseRule(&request.Dummy)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(dummy *dummy.Dummy) {
	err := s.repository.Edit(dummy)
	if err != nil {
		s.Error = &services.Error{
			Status:  400,
			Message: "Try again in a few minutes",
			Error:   "Error on request process",
		}
		return
	}
	s.response = &Response{
		Status: http.StatusOK,
	}
}
