package handlers

import (
	"go-skeleton/internal/application/domain/user"
	"go-skeleton/internal/application/providers/pagination"
	userCreate "go-skeleton/internal/application/services/user/CREATE"
	userDelete "go-skeleton/internal/application/services/user/DELETE"
	userEdit "go-skeleton/internal/application/services/user/EDIT"
	userGet "go-skeleton/internal/application/services/user/GET"
	userList "go-skeleton/internal/application/services/user/LIST"
	userRepository "go-skeleton/internal/repositories/user"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
	"go-skeleton/pkg/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandlers struct {
	UserRepository *userRepository.UserRepository

	logger    *logger.Logger
	idCreator *idCreator.IdCreator
	validator *validator.Validator
}

func NewUserHandlers(reg *registry.Registry) *UserHandlers {
	return &UserHandlers{
		UserRepository: reg.Inject("userRepository").(*userRepository.UserRepository),
		logger:          reg.Inject("logger").(*logger.Logger),
		idCreator:       reg.Inject("idCreator").(*idCreator.IdCreator),
		validator:       reg.Inject("validator").(*validator.Validator),
	}
}

func (hs *UserHandlers) HandleGetUser(context echo.Context) error {
	s := userGet.NewService(hs.logger, hs.UserRepository)
	data := new(userGet.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		userGet.NewRequest(data),
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}

func (hs *UserHandlers) HandleCreateUser(context echo.Context) error {
	s := userCreate.NewService(hs.logger, hs.UserRepository, hs.idCreator)
	data := new(userCreate.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(http.StatusBadRequest, errors)
	}

	s.Execute(
		userCreate.NewRequest(data, hs.validator),
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusCreated, response)
}

func (hs *UserHandlers) HandleEditUser(context echo.Context) error {
	s := userEdit.NewService(hs.logger, hs.UserRepository)
	data := new(userEdit.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(http.StatusBadRequest, errors)
	}

	s.Execute(
		userEdit.NewRequest(data, hs.validator),
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}

func (hs *UserHandlers) HandleListUser(context echo.Context) error {
	s := userList.NewService(
		hs.logger,
		hs.UserRepository,
		pagination.NewPaginationProvider[user.User](hs.UserRepository),
	)

	data := new(userList.Data)
	bindErr := echo.QueryParamsBinder(context).
		Int("page", &data.Page).
		BindErrors()

	if bindErr != nil {
		return context.JSON(http.StatusBadRequest, bindErr)
	}

	s.Execute(
		userList.NewRequest(data),
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}

func (hs *UserHandlers) HandleDeleteUser(context echo.Context) error {
	s := userDelete.NewService(hs.logger, hs.UserRepository)
	data := new(userDelete.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(http.StatusBadRequest, errors)
	}

	s.Execute(
		userDelete.NewRequest(data),
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}
