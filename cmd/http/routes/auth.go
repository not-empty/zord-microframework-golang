package routes

import (
	"github.com/labstack/echo/v4"
	"go-skeleton/cmd/http/handlers"
)

type Auth struct {
	hand handlers.AuthHandlers
}

func NewAuthRoute() *Auth {
	return &Auth{}
}

func (hs *Auth) DeclareRoutes(server *echo.Group) {
	server.POST("/login", hs.hand.HandleLogin)
}
