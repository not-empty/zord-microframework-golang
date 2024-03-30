package routes

import (
	"go-skeleton/cmd/http/handlers"
	"go-skeleton/pkg/registry"

	"github.com/labstack/echo/v4"
)

type DummyRoutes struct {
	hand *handlers.DummyHandlers
}

func NewDummyRoutes(reg *registry.Registry) *DummyRoutes {
	hand := handlers.NewDummyHandlers(reg)
	return &DummyRoutes{
		hand: hand,
	}
}

func (hs *DummyRoutes) DeclareRoutes(server *echo.Group) {
	server.GET("/v1/dummy", hs.hand.HandleListDummy)
	server.GET("/v1/dummy/:id", hs.hand.HandleGetDummy)
	server.POST("/v1/dummy", hs.hand.HandleCreateDummy)
	server.PUT("/v1/dummy/:id", hs.hand.HandleEditDummy)
	server.DELETE("/v1/dummy/:id", hs.hand.HandleDeleteDummy)
}
