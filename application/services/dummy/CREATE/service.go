package dummy

import (
	"encoding/json"
	"fmt"
	"go-skeleton/application/domain/dummy"
	"go-skeleton/application/services"
	"io"
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
	s.Logger.Debug("Creating new dummy")

	body, err := io.ReadAll(request.Body)
	if err != nil {
		s.BadRequest(request, err)
		return
	}

	err = json.Unmarshal(body, &request.Dummy)
	if err != nil {
		s.BadRequest(request, err)
		return
	}

	if err := request.Validate(); err != nil {
		s.BadRequest(request, err)
		return
	}

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
