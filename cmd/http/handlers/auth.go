package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-skeleton/application/domain/auth"
	login "go-skeleton/application/services/auth/LOGIN"
	"go-skeleton/pkg"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"
	"io"
)

type AuthHandlers struct {
	logger        *logger.Logger
	Secret        string
	JwtExpiration int
	AccessSecret  []string
	AccessContext []string
	AccessToken   []string
}

func NewAuthHandlers(deps map[string]pkg.Bootable) *AuthHandlers {
	config := deps["config"].(*config.Config)

	return &AuthHandlers{
		logger:        deps["logger"].(*logger.Logger),
		Secret:        config.ReadConfig("JWT_SECRET"),
		JwtExpiration: config.ReadNumberConfig("JWT_EXPIRATION"),
		AccessSecret:  config.ReadArrayConfig("ACCESS_SECRET"),
		AccessContext: config.ReadArrayConfig("ACCESS_CONTEXT"),
		AccessToken:   config.ReadArrayConfig("ACCESS_TOKEN"),
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

	s := login.NewService(
		hs.logger,
		hs.Secret,
		hs.JwtExpiration,
		hs.AccessSecret,
		hs.AccessContext,
		hs.AccessToken,
	)
	s.Execute(
		login.NewRequest(domain),
	)

	response, err := s.GetResponse()
	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(response.Status, response)
}
