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

func (s *Service) Execute(request Request) {
	if err := request.Validate(); err != nil {
		s.BadRequest(err.Error())
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
		ID:        data.DummyId,
		DummyName: data.DummyName,
	}

	tx, txErr := s.repository.InitTX()
	if txErr != nil {
		s.InternalServerError("error on create", txErr)
		return
	}

	err := s.repository.Create(dummy, tx, true)
	if err != nil {
		s.InternalServerError("error on create", err)
		return
	}

	s.response = &Response{
		Data: dummy,
	}
}
