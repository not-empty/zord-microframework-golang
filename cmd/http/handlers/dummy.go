package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-skeleton/application/domain/dummy"
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
	"io"
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
	s.Execute(
		dummyGet.NewRequest(context.Param("dummy_id")),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}

func (hs *DummyHandlers) HandleCreateDummy(context echo.Context) error {
	s := dummyCreate.NewService(hs.logger, hs.DummyRepository, hs.idCreator)

	domain := dummy.Dummy{}
	body, errors := io.ReadAll(context.Request().Body)
	if errors != nil {
		return context.JSON(422, errors)
	}

	errors = json.Unmarshal(body, &domain)
	if errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		dummyCreate.NewRequest(domain, hs.validator),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}

func (hs *DummyHandlers) HandleEditDummy(context echo.Context) error {
	s := dummyEdit.NewService(hs.logger, hs.DummyRepository)

	domain := dummy.Dummy{}
	body, errors := io.ReadAll(context.Request().Body)
	if errors != nil {
		return context.JSON(422, errors)
	}

	errors = json.Unmarshal(body, &domain)
	if errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		dummyEdit.NewRequest(domain, context.Param("dummy_id"), hs.validator),
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
	s.Execute(
		dummyDelete.NewRequest(context.Param("dummy_id")),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}
