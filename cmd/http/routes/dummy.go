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

func (hs *DummyRoutes) DeclarePrivateRoutes(server *echo.Group, apiPrefix string) {
	server.GET(apiPrefix+"/dummy", hs.hand.HandleListDummy)
	server.GET(apiPrefix+"/dummy/:id", hs.hand.HandleGetDummy)
	server.POST(apiPrefix+"/dummy", hs.hand.HandleCreateDummy)
	server.PUT(apiPrefix+"/dummy/:id", hs.hand.HandleEditDummy)
	server.DELETE(apiPrefix+"/dummy/:id", hs.hand.HandleDeleteDummy)
}

func (hs *DummyRoutes) DeclarePublicRoutes(server *echo.Group, apiPrefix string) {}
