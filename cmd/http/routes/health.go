package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Health struct {
}

func NewHealthRoute() *Health {
	return &Health{}
}

func (hs *Health) DeclarePrivateRoutes(server *echo.Group, apiPrefix string) {}

func (hs *Health) DeclarePublicRoutes(server *echo.Group, apiPrefix string) {
	server.GET(apiPrefix+"/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})
}
