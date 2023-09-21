package dependencies

import (
	"go-skeleton/cmd/handlers/http/middlewares"
	"go-skeleton/pkg/config"

	"github.com/labstack/echo/v4"
)

type Middleware interface {
	GetMiddleware() echo.MiddlewareFunc
	Boot(*config.Config)
}

var auth = middlewares.NewJwtAuth()

var MiddlewareList = map[string]Middleware{
	"auth": auth,
}
