package handlers

import (
	dummyCreate "go-skeleton/application/services/dummy/CREATE"
	dummyDelete "go-skeleton/application/services/dummy/DELETE"
	dummyEdit "go-skeleton/application/services/dummy/EDIT"
	dummyGet "go-skeleton/application/services/dummy/GET"
	dummyList "go-skeleton/application/services/dummy/LIST"
	"go-skeleton/pkg"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	dummyRepository "go-skeleton/pkg/repositories/dummy"
	"go-skeleton/pkg/validator"

	"github.com/labstack/echo/v4"
)

type DummyHandlers struct {
	Environment     string
	DummyRepository *dummyRepository.DummyRepository

	logger    *logger.Logger
	idCreator *idCreator.IdCreator
	validator *validator.Validator
}

func NewDummyHandlers(
	environment string,
	deps map[string]pkg.Bootable,
) *DummyHandlers {
	l := deps["logger"].(*logger.Logger)
	m := deps["mysql"].(*database.MySql)
	i := deps["IdCreator"].(*idCreator.IdCreator)
	v := deps["validator"].(*validator.Validator)
	repository := dummyRepository.NewBaseRepository(m)

	return &DummyHandlers{
		Environment:     environment,
		DummyRepository: repository,
		logger:          l,
		idCreator:       i,
		validator:       v,
	}
}

func (hs *DummyHandlers) HandleGetDummy(context echo.Context) error {
	s := dummyGet.NewService(hs.logger, hs.DummyRepository)
	dto := new(dummyGet.RequestDTO)

	if errors := context.Bind(dto); errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		dummyGet.NewRequest(dto),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}

func (hs *DummyHandlers) HandleCreateDummy(context echo.Context) error {
	s := dummyCreate.NewService(hs.logger, hs.DummyRepository, hs.idCreator)
	dto := new(dummyCreate.RequestDTO)

	if errors := context.Bind(dto); errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		dummyCreate.NewRequest(dto, hs.validator),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}

func (hs *DummyHandlers) HandleEditDummy(context echo.Context) error {
	s := dummyEdit.NewService(hs.logger, hs.DummyRepository)
	dto := new(dummyEdit.RequestDTO)

	if errors := context.Bind(dto); errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		dummyEdit.NewRequest(dto, hs.validator),
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
	dto := new(dummyDelete.RequestDTO)

	if errors := context.Bind(dto); errors != nil {
		return context.JSON(422, errors)
	}

	s.Execute(
		dummyDelete.NewRequest(dto),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}
