package handlers

import (
	"github.com/labstack/echo/v4"
	userDomain "go-skeleton/internal/application/domain/user"
	"go-skeleton/internal/application/services/user"
	user2 "go-skeleton/internal/repositories/user"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
)

type UserHandlers struct {
	logger *logger.Logger
	repo   userDomain.Repo
}

func NewUserHandlers(reg *registry.Registry) *UserHandlers {
	return &UserHandlers{
		logger: reg.Inject("logger").(*logger.Logger),
		repo:   reg.Inject("userRepo").(*user2.Repository),
	}
}

func (hs *UserHandlers) HandleUser(context echo.Context) error {
	s := user.NewService(hs.logger, hs.repo)

	data := new(user.Data)
	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		user.NewRequest(data),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(422, err)
	}
	return context.JSON(200, response)
}
