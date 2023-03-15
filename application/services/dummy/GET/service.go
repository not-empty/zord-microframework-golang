package dummy

import (
	"go-skeleton/application/services"
	"net/http"
)

type Service struct {
	services.BaseService
	response *Response
}

func NewService(log services.Logger) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
	}
}

func (s *Service) Execute(request Request) {
	s.Logger.Debug("Hello Im Dummy Server!")
	if err := request.Validate(); err != nil {
		s.BadRequest(request, err)
		return
	}
	s.produceResponseRule()
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule() {
	s.Logger.Debug("ProduceResponseRule")
	if s.Error == nil {
		s.response = &Response{
			Status:  http.StatusOK,
			Message: "OK",
		}
	}
}
