package handlers

import (
	"encoding/json"
	"go-skeleton/internal/application/domain/auth"
	authLogin "go-skeleton/internal/application/services/auth/LOGIN"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
	"io"

	"github.com/labstack/echo/v4"
)

type AuthHandlers struct {
	logger        *logger.Logger
	Secret        string
	JwtExpiration int
	AccessSecret  []string
	AccessContext []string
	AccessToken   []string
}

func NewAuthHandlers(reg *registry.Registry) *AuthHandlers {
	conf := reg.Inject("config").(*config.Config)
	return &AuthHandlers{
		logger:        reg.Inject("logger").(*logger.Logger),
		Secret:        conf.ReadConfig("JWT_SECRET"),
		JwtExpiration: conf.ReadNumberConfig("JWT_EXPIRATION"),
		AccessSecret:  conf.ReadArrayConfig("ACCESS_SECRET"),
		AccessContext: conf.ReadArrayConfig("ACCESS_CONTEXT"),
		AccessToken:   conf.ReadArrayConfig("ACCESS_TOKEN"),
	}
}

func (hs *AuthHandlers) HandleLogin(c echo.Context) error {
	domain := auth.Token{}

	body, errors := io.ReadAll(c.Request().Body)
	if errors != nil {
		return c.JSON(422, errors)
	}

	errors = json.Unmarshal(body, &domain)
	if errors != nil {
		return c.JSON(422, errors)
	}

	s := authLogin.NewService(
		hs.logger,
		hs.Secret,
		hs.JwtExpiration,
		hs.AccessSecret,
		hs.AccessContext,
		hs.AccessToken,
	)

	s.Execute(
		authLogin.NewRequest(domain),
	)

	response, err := s.GetResponse()
	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(response.Status, response)
}
