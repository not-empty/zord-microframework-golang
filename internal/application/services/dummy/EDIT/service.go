package dummy

import (
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/application/services"
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
	if err := request.Validate(); err != nil {
		s.BadRequest(err.Error())
		return
	}

	s.produceResponseRule(request.Data)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(data *Data) {
	dummyDomain := dummy.Dummy{
		ID:        data.ID,
		DummyName: data.DummyName,
	}

	affected, err := s.repository.Edit(dummyDomain, "id", data.ID)
	if err != nil {
		s.InternalServerError("error on edit")
		return
	}

	if affected < 1 {
		s.UnprocessableEntity("same data or invalid id")
		return
	}

	s.response = &Response{
		Data: data,
	}
}
