package handlers

import (
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/application/providers/pagination"
	dummyCreate "go-skeleton/internal/application/services/dummy/CREATE"
	dummyDelete "go-skeleton/internal/application/services/dummy/DELETE"
	dummyEdit "go-skeleton/internal/application/services/dummy/EDIT"
	dummyGet "go-skeleton/internal/application/services/dummy/GET"
	dummyList "go-skeleton/internal/application/services/dummy/LIST"
	dummyRepository "go-skeleton/internal/repositories/dummy"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
	"go-skeleton/pkg/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DummyHandlers struct {
	DummyRepository *dummyRepository.DummyRepository

	logger    *logger.Logger
	idCreator *idCreator.IdCreator
	validator *validator.Validator
}

func NewDummyHandlers(reg *registry.Registry) *DummyHandlers {
	return &DummyHandlers{
		DummyRepository: reg.Inject("dummyRepository").(*dummyRepository.DummyRepository),
		logger:          reg.Inject("logger").(*logger.Logger),
		idCreator:       reg.Inject("idCreator").(*idCreator.IdCreator),
		validator:       reg.Inject("validator").(*validator.Validator),
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
	return context.JSON(http.StatusOK, response)
}

func (hs *DummyHandlers) HandleCreateDummy(context echo.Context) error {
	s := dummyCreate.NewService(hs.logger, hs.DummyRepository, hs.idCreator)
	data := new(dummyCreate.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(http.StatusBadRequest, errors)
	}

	s.Execute(
		dummyCreate.NewRequest(data, hs.validator),
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusCreated, response)
}

func (hs *DummyHandlers) HandleEditDummy(context echo.Context) error {
	s := dummyEdit.NewService(hs.logger, hs.DummyRepository)
	data := new(dummyEdit.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(http.StatusBadRequest, errors)
	}

	s.Execute(
		dummyEdit.NewRequest(data, hs.validator),
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}

func (hs *DummyHandlers) HandleListDummy(context echo.Context) error {
	s := dummyList.NewService(
		hs.logger,
		hs.DummyRepository,
		pagination.NewPaginationProvider[dummy.Dummy](hs.DummyRepository),
	)

	data := new(dummyList.Data)
	bindErr := echo.QueryParamsBinder(context).
		Int("page", &data.Page).
		BindErrors()

	if bindErr != nil {
		return context.JSON(http.StatusBadRequest, bindErr)
	}

	s.Execute(
		dummyList.NewRequest(data),
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}

func (hs *DummyHandlers) HandleDeleteDummy(context echo.Context) error {
	s := dummyDelete.NewService(hs.logger, hs.DummyRepository)
	data := new(dummyDelete.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(http.StatusBadRequest, errors)
	}

	s.Execute(
		dummyDelete.NewRequest(data),
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}
