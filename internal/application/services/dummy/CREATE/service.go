package dummy

import (
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/application/services"
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

	request.Data.DummyId = s.idCreator.Create()
	s.produceResponseRule(request.Data)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(data *Data) {
	dummy := dummy.Dummy{
		DummyId:   data.DummyId,
		DummyName: data.DummyName,
	}
	err := s.repository.Create(&dummy)
	if err != nil {
		s.Error = &services.Error{
			Status:  http.StatusInternalServerError,
			Message: "Try again in a few minutes",
			Error:   "Error on request process",
		}
		return
	}
	s.response = &Response{
		Data: dummy,
	}
}
