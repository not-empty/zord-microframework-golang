package dummy

import (
	"fmt"
	"go-skeleton/application/domain/dummy"
	"go-skeleton/application/services"
	"net/http"
)

type Service struct {
	services.BaseService
	response   *Response
	repository dummy.Repository
	idCreator  services.IdCreator
}

func NewService(log services.Logger, repository dummy.Repository, idCreator services.IdCreator) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
		repository: repository,
		idCreator:  idCreator,
	}
}

func (s *Service) Execute(request Request) {
	s.Logger.Debug("Creating new dummy")

	if err := request.Validate(); err != nil {
		s.BadRequest(request, err)
		return
	}

	request.Dummy.DummyId = s.idCreator.Create()
	s.produceResponseRule(request.Dummy)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(dummy dummy.Dummy) {
	status := s.repository.Create(&dummy)
	if s.Error == nil {
		s.response = &Response{
			Status:  http.StatusOK,
			Message: fmt.Sprint(status),
		}
	}
}
