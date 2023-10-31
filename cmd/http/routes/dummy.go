package routes

import (
	"github.com/labstack/echo/v4"
	"go-skeleton/cmd/http/handlers"
)

type DummyRoutes struct {
	hand *handlers.DummyHandlers
}

func NewDummyRoutes() *DummyRoutes {
	hand := handlers.NewDummyHandlers()
	return &DummyRoutes{
		hand: hand,
	}
}

func (hs *DummyRoutes) DeclareRoutes(server *echo.Group) {
	server.GET("/v1/dummy", hs.hand.HandleListDummy)
	server.GET("/v1/dummy/:dummy_id", hs.hand.HandleGetDummy)
	server.POST("/v1/dummy", hs.hand.HandleCreateDummy)
	server.PUT("/v1/dummy/:dummy_id", hs.hand.HandleEditDummy)
	server.DELETE("/v1/dummy/:dummy_id", hs.hand.HandleDeleteDummy)
}
