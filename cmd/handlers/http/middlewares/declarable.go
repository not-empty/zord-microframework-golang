package middlewares

import (
	"github.com/labstack/echo/v4"
)

type Middleware interface {
	GetMiddleware() echo.MiddlewareFunc
	Boot()
}

func GetAllMiddlewares(secret string) map[string]Middleware {
	auth := NewJwtAuth(secret)
	return map[string]Middleware{
		"auth": auth,
	}
}
