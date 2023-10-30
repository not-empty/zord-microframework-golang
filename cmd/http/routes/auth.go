package routes

import (
	"go-skeleton/cmd/http/handlers"
	"go-skeleton/pkg"

	"github.com/labstack/echo/v4"
)

type Auth struct {
	hand *handlers.AuthHandlers
}

func NewAuthRoute(deps map[string]pkg.Bootable) *Auth {
	return &Auth{
		hand: handlers.NewAuthHandlers(deps),
	}
}

func (hs *Auth) DeclareRoutes(server *echo.Group) {
	server.POST("/login", hs.hand.HandleLogin)
}
