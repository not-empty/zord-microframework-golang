package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Health struct {
}

func NewHealthRoute() *Health {
	return &Health{}
}

func (hs *Health) DeclareRoutes(server *echo.Group) {
	server.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})
}
