package middlewares

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(secret string) echo.MiddlewareFunc {
	Config := echojwt.Config{
		SigningKey: []byte(secret),
	}

	return echojwt.WithConfig(Config)
}
