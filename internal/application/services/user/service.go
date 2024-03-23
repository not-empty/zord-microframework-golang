package user

import (
	"context"
	userDomain "go-skeleton/internal/application/domain/user"
	"go-skeleton/internal/application/services"
)

type Service struct {
	services.BaseService
	response *Response
	Repo     userDomain.Repo
}

func NewService(log services.Logger, repo userDomain.Repo) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
		Repo: repo,
	}
}

func (s *Service) Execute(request Request) {

	if err := request.Validate(); err != nil {
		s.BadRequest(request, err)
		return
	}

	s.produceResponseRule(request.Data)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(data *Data) {
	err, u := s.Repo.Create(context.Background(), userDomain.User{
		Name:     "Samuel da Silva",
		Email:    "samuel2@gmail.com",
		Password: "1232323234343",
	})

	if err != nil {
		return
	}

	s.response = &Response{
		Data: userDomain.User{
			Id:       u.Id,
			Name:     u.Name,
			Email:    u.Email,
			Password: u.Password,
		},
	}
}
