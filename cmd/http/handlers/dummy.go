package handlers

import (
	dummyCreate "go-skeleton/application/services/dummy/CREATE"
	dummyDelete "go-skeleton/application/services/dummy/DELETE"
	dummyEdit "go-skeleton/application/services/dummy/EDIT"
	dummyGet "go-skeleton/application/services/dummy/GET"
	dummyList "go-skeleton/application/services/dummy/LIST"
	"go-skeleton/pkg"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	dummyRepository "go-skeleton/pkg/repositories/dummy"
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
	s := dummyGet.NewService(hs.logger, hs.DummyRepository)
	data := new(dummyGet.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		dummyGet.NewRequest(data),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}

func (hs *DummyHandlers) HandleCreateDummy(context echo.Context) error {
	s := dummyCreate.NewService(hs.logger, hs.DummyRepository, hs.idCreator)
	data := new(dummyCreate.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		dummyCreate.NewRequest(data, hs.validator),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}

func (hs *DummyHandlers) HandleEditDummy(context echo.Context) error {
	s := dummyEdit.NewService(hs.logger, hs.DummyRepository)
	data := new(dummyEdit.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		dummyEdit.NewRequest(data, hs.validator),
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}

func (hs *DummyHandlers) HandleListDummy(context echo.Context) error {
	s := dummyList.NewService(hs.logger, hs.DummyRepository)
	s.Execute(
		dummyList.NewRequest(),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}

func (hs *DummyHandlers) HandleDeleteDummy(context echo.Context) error {
	s := dummyDelete.NewService(hs.logger, hs.DummyRepository)
	data := new(dummyDelete.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		dummyDelete.NewRequest(data),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}
