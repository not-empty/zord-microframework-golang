package dummy

import (
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/application/services"
	"math"
	"net/http"
)

const ListLimit = 25.0

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
	s.produceResponseRule(request.Data.Page)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(page int) {
	dummyData := &[]dummy.Dummy{}
	offset := (page - 1) * ListLimit

	total, err := s.repository.Count()
	if err != nil {
		s.Error = &services.Error{
			Status:  http.StatusInternalServerError,
			Message: "Try again in a few minutes",
			Error:   "fatal error",
		}
		return
	}

	if total == 0 {
		s.Error = &services.Error{
			Status:  http.StatusNotFound,
			Message: "Try again in a few minutes",
			Error:   "data not found",
		}
		return
	}

	totalPages := math.Ceil(float64(total) / ListLimit)
	if page <= int(totalPages) {
		dummyData, err = s.repository.List(ListLimit, offset)
		if err != nil {
			s.Error = &services.Error{
				Status:  http.StatusInternalServerError,
				Message: "Try again in a few minutes",
				Error:   "Error on request process",
			}
			return
		}
	}

	s.response = &Response{
		Status:      http.StatusOK,
		CurrentPage: page,
		TotalPages:  int64(totalPages),
		Data:        dummyData,
	}
}
