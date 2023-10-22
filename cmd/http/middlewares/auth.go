package middlewares

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JwtAuth struct {
	Secret string
}

func NewJwtAuth(Secret string) *JwtAuth {
	return &JwtAuth{
		Secret: Secret,
	}
}

func (j *JwtAuth) Boot() {
}

func (j *JwtAuth) GetMiddleware() echo.MiddlewareFunc {
	Config := echojwt.Config{
		SigningKey: []byte(j.Secret),
	}

	return echojwt.WithConfig(Config)
}
