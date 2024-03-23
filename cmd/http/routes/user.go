package routes

import (
	"github.com/labstack/echo/v4"
	"go-skeleton/cmd/http/handlers"
	"go-skeleton/pkg/registry"
)

type UserRoutes struct {
	hand *handlers.UserHandlers
}

func NewUserRoutes(reg *registry.Registry) *UserRoutes {
	hand := handlers.NewUserHandlers(reg)
	return &UserRoutes{
		hand: hand,
	}
}

func (hs *UserRoutes) DeclareRoutes(server *echo.Group) {
	server.GET("/v1/user", hs.hand.HandleUser)
}
