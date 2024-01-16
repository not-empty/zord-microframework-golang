package routes

import (
	"go-skeleton/cmd/http/handlers"
	"go-skeleton/pkg/registry"

	"github.com/labstack/echo/v4"
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
	server.GET("/v1/user", hs.hand.HandleListUser)
	server.GET("/v1/user/:user_id", hs.hand.HandleGetUser)
	server.POST("/v1/user", hs.hand.HandleCreateUser)
	server.PUT("/v1/user/:user_id", hs.hand.HandleEditUser)
	server.DELETE("/v1/user/:user_id", hs.hand.HandleDeleteUser)
}
