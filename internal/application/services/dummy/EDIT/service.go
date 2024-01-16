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

	affectedRows, err := s.repository.Edit(&dummy)
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
		*data,
	}
}
