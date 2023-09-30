package routes

import (
	"encoding/json"
	"go-skeleton/application/domain/dummy"
	dummyCreate "go-skeleton/application/services/dummy/CREATE"
	dummyDelete "go-skeleton/application/services/dummy/DELETE"
	dummyEdit "go-skeleton/application/services/dummy/EDIT"
	dummyGet "go-skeleton/application/services/dummy/GET"
	dummyList "go-skeleton/application/services/dummy/LIST"
	"go-skeleton/cmd/handlers/types"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	dummyRepository "go-skeleton/pkg/repositories/dummy"
	"go-skeleton/pkg/validator"
	"io"

	"github.com/labstack/echo/v4"
)

type DummyRoutes struct {
	Environment     string
	DummyRepository *dummyRepository.DummyRepository

	logger    *logger.Logger
	idCreator *idCreator.IdCreator
	validator *validator.Validator
}

func NewDummyRoutes(server *types.Server) *DummyRoutes {
	repository := dummyRepository.NewBaseRepository(server.Mysql)

	return &DummyRoutes{
		logger:          server.Logger,
		Environment:     server.Environment,
		DummyRepository: repository,
		idCreator:       server.IdCreator,
		validator:       server.Validator,
	}
}

func (hs *DummyRoutes) DeclareRoutes(server *echo.Group) {
	server.GET("/v1/dummy", hs.HandleListDummy)
	server.GET("/v1/dummy/:dummy_id", hs.HandleGetDummy)
	server.POST("/v1/dummy", hs.HandleCreateDummy)
	server.PUT("/v1/dummy/:dummy_id", hs.HandleEditDummy)
	server.DELETE("/v1/dummy/:dummy_id", hs.HandleDeleteDummy)
}

func (hs *DummyRoutes) HandleGetDummy(context echo.Context) error {
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

func (hs *DummyRoutes) HandleCreateDummy(context echo.Context) error {
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

func (hs *DummyRoutes) HandleEditDummy(context echo.Context) error {
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

func (hs *DummyRoutes) HandleListDummy(context echo.Context) error {
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

func (hs *DummyRoutes) HandleDeleteDummy(context echo.Context) error {
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
