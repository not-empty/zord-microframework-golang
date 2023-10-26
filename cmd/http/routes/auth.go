package routes

import (
	"github.com/labstack/echo/v4"
	"go-skeleton/cmd/http/handlers"
	"go-skeleton/pkg"
)

type Auth struct {
	hand handlers.AuthHandlers
}

func NewAuthRoute(
	deps map[string]pkg.Bootable,
) *Auth {
	return &Auth{}
}

func (hs *Auth) DeclareRoutes(server *echo.Group) {
	server.POST("/login", hs.hand.HandleLogin)
}
