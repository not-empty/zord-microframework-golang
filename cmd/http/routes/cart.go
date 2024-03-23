package routes

import (
	"github.com/labstack/echo/v4"
	"go-skeleton/cmd/http/handlers"
	"go-skeleton/pkg/registry"
)

type CartRoutes struct {
	hand *handlers.CartHandlers
}

func NewCartRoutes(reg *registry.Registry) *CartRoutes {
	hand := handlers.NewCartHandlers(reg)
	return &CartRoutes{
		hand: hand,
	}
}

func (hs *CartRoutes) DeclareRoutes(server *echo.Group) {
	server.GET("/v1/cart", hs.hand.HandleCart)
}
