package routes

import (
	"go-skeleton/application/domain/dummy"
	dummyCreate "go-skeleton/application/services/dummy/CREATE"
	dummyGet "go-skeleton/application/services/dummy/GET"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/repositories"

	"github.com/labstack/echo/v4"
)

type DummyRoutes struct {
	Environment     string
	DummyRepository *repositories.BaseRepository[dummy.Dummy]

	// config *config.Config
	logger *logger.Logger
}

func NewDummyRoutes(logger *logger.Logger, Environment string) *DummyRoutes {
	repository := &repositories.BaseRepository[dummy.Dummy]{}
	return &DummyRoutes{
		logger:          logger,
		Environment:     Environment,
		DummyRepository: repository,
	}
}

func (hs *DummyRoutes) DeclareRoutes(server *echo.Echo) {
	server.GET("/v1/dummy/:dummy_id", hs.HandleGetDummy)
	server.POST("/v1/dummy", hs.HandleCreateDummy)
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
	s := dummyCreate.NewService(hs.logger, hs.DummyRepository)
	s.Execute(
		dummyCreate.NewRequest(),
	)
	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(response.Status, response)
}
