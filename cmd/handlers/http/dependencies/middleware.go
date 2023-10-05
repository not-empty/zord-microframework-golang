package dependencies

import (
	"github.com/labstack/echo/v4"
	"go-skeleton/cmd/handlers/http/middlewares"
)

type Middleware interface {
	GetMiddleware() echo.MiddlewareFunc
	Boot()
}

func GetAllMiddlewares(secret string) map[string]Middleware {
	auth := middlewares.NewJwtAuth(secret)
	return map[string]Middleware{
		"auth": auth,
	}
}
