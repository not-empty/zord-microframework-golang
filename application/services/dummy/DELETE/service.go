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
	repository services.Repository[dummy.Dummy]
}

func NewService(log services.Logger, repository services.Repository[dummy.Dummy]) *Service {
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
	s.produceResponseRule(request.Dummy.DummyId)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(id string) {
	s.Logger.Debug("ProduceResponseRule")
	teste := s.repository.Delete(id)
	if s.Error == nil {
		s.response = &Response{
			Status:  http.StatusOK,
			Message: fmt.Sprint(teste),
		}
	}
}
