package dummy

import (
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/application/services"
)

type Service struct {
	services.BaseService
	response   *Response
	repository dummy.Repository
	pagProv    dummy.PaginationProvider
}

func NewService(log services.Logger, repository dummy.Repository, pagProv dummy.PaginationProvider) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
		repository: repository,
		pagProv:    pagProv,
	}
}

func (s *Service) Execute(request Request) {
	if err := request.Validate(); err != nil {
		s.BadRequest(err.Error())
		return
	}
	if err := request.SetFiltersRules(); err != nil {
		s.BadRequest(err.Error())
		return
	}
	request.Domain.SetFilters(request.Filters)

	s.produceResponseRule(request.Data.Page, 25, request.Domain)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(page int, limit int, domain *dummy.Dummy) {
	err, pagination := s.pagProv.PaginationHandler(*domain, page, limit)
	if err != nil {
		s.CustomError(err.Status, err)
		return
	}

	s.response = &Response{
		CurrentPage: page,
		TotalPages:  pagination.TotalPages,
		Data:        pagination.Data,
	}
}
