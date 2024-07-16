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

	s.produceResponseRule(request.ID, request.Client, request.Data)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(id string, client string, data *Data) {
	dummyDomain := dummy.Dummy{
		ID:        id,
		DummyName: data.DummyName,
		Email:     data.Email,
	}
	dummyDomain = dummyDomain.SetClient(client)

	affected, err := s.repository.Edit(dummyDomain, "id", id)
	if err != nil {
		s.InternalServerError("error on edit", err)
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
