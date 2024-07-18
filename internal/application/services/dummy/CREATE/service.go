package dummy

import (
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/application/services"
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

func (s *Service) Execute(request *Request) {
	if err := request.Validate(); err != nil {
		s.BadRequest(err.Error())
		return
	}
	s.produceResponseRule(request.Data, request.Domain)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(data *Data, dummyData *dummy.Dummy) {
	dummyData.ID = s.idCreator.Create()
	dummyData.DummyName = data.DummyName
	dummyData.Email = data.Email

	tx, txErr := s.repository.InitTX()
	if txErr != nil {
		s.InternalServerError("error on create", txErr)
		return
	}

	err := s.repository.Create(*dummyData, tx, true)
	if err != nil {
		s.InternalServerError("error on create", err)
		return
	}

	s.response = &Response{
		Data: *dummyData,
	}
}
