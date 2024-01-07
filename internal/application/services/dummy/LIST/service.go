package dummy

import (
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/application/providers/pagination"
	"go-skeleton/internal/application/services"
	"net/http"
)

type Service struct {
	services.BaseService
	response   *Response
	repository dummy.Repository
	limit      int
}

func NewService(log services.Logger, repository dummy.Repository, limit int) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
		repository: repository,
		limit:      limit,
	}
}

func (s *Service) Execute(request Request) {
	s.Logger.Debug("Hello Im Dummy Server!")
	if err := request.Validate(); err != nil {
		s.BadRequest(request, err)
		return
	}
	s.produceResponseRule(request.Data.Page)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(page int) {
	err, pagination := pagination.PaginationProvider[dummy.Dummy](s.repository, page, s.limit)
	if err != nil {
		s.Error = err
		return
	}

	s.response = &Response{
		Status:      http.StatusOK,
		CurrentPage: page,
		TotalPages:  pagination.TotalPages,
		Data:        pagination.Data,
	}
}
