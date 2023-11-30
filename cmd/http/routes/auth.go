package routes

import (
	"go-skeleton/cmd/http/handlers"
	"go-skeleton/pkg/registry"

	"github.com/labstack/echo/v4"
)

type Auth struct {
	hand *handlers.AuthHandlers
}

func NewAuthRoute(reg *registry.Registry) *Auth {
	return &Auth{
		hand: handlers.NewAuthHandlers(reg),
	}
}

func (hs *Auth) DeclareRoutes(server *echo.Group) {
	server.POST("/login", hs.hand.HandleLogin)
}
