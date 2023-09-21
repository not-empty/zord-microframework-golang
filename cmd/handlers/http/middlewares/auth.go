package middlewares

import (
	"go-skeleton/pkg/config"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JwtAuth struct {
	Secret string
}

func NewJwtAuth() *JwtAuth {
	return &JwtAuth{}
}

func (j *JwtAuth) Boot(conf *config.Config) {
	j.Secret = conf.Secret
}

func (j *JwtAuth) GetMiddleware() echo.MiddlewareFunc {
	Config := echojwt.Config{
		SigningKey: []byte(j.Secret),
	}

	return echojwt.WithConfig(Config)
}
