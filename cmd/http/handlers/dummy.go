package handlers

import (
	dummy5 "go-skeleton/internal/application/services/dummy/CREATE"
	dummy4 "go-skeleton/internal/application/services/dummy/DELETE"
	dummy2 "go-skeleton/internal/application/services/dummy/EDIT"
	"go-skeleton/internal/application/services/dummy/GET"
	dummy3 "go-skeleton/internal/application/services/dummy/LIST"
	"go-skeleton/internal/repositories/dummy"
	"go-skeleton/pkg"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/validator"

	"github.com/labstack/echo/v4"
)

type DummyHandlers struct {
	DummyRepository *dummyRepository.DummyRepository

	logger    *logger.Logger
	idCreator *idCreator.IdCreator
	validator *validator.Validator
}

func NewDummyHandlers() *DummyHandlers {
	repository := dummyRepository.NewBaseRepository(pkg.Mysql)

	return &DummyHandlers{
		DummyRepository: repository,
		logger:          pkg.Logger,
		idCreator:       pkg.IdCreator,
		validator:       pkg.Validator,
	}
}

func (hs *DummyHandlers) HandleGetDummy(context echo.Context) error {
	s := dummy.NewService(hs.logger, hs.DummyRepository)
	data := new(dummy.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		dummy.NewRequest(data),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}

func (hs *DummyHandlers) HandleCreateDummy(context echo.Context) error {
	s := dummy5.NewService(hs.logger, hs.DummyRepository, hs.idCreator)
	data := new(dummy5.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		dummy5.NewRequest(data, hs.validator),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}

func (hs *DummyHandlers) HandleEditDummy(context echo.Context) error {
	s := dummy2.NewService(hs.logger, hs.DummyRepository)
	data := new(dummy2.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		dummy2.NewRequest(data, hs.validator),
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}

func (hs *DummyHandlers) HandleListDummy(context echo.Context) error {
	s := dummy3.NewService(hs.logger, hs.DummyRepository)
	s.Execute(
		dummy3.NewRequest(),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}

func (hs *DummyHandlers) HandleDeleteDummy(context echo.Context) error {
	s := dummy4.NewService(hs.logger, hs.DummyRepository)
	data := new(dummy4.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		dummy4.NewRequest(data),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}
