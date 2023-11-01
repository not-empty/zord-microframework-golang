package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-skeleton/internal/application/domain/auth"
	auth2 "go-skeleton/internal/application/services/auth/LOGIN"
	"go-skeleton/pkg"
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

func NewAuthHandlers() *AuthHandlers {
	return &AuthHandlers{
		logger:        pkg.Logger,
		Secret:        pkg.Config.ReadConfig("JWT_SECRET"),
		JwtExpiration: pkg.Config.ReadNumberConfig("JWT_EXPIRATION"),
		AccessSecret:  pkg.Config.ReadArrayConfig("ACCESS_SECRET"),
		AccessContext: pkg.Config.ReadArrayConfig("ACCESS_CONTEXT"),
		AccessToken:   pkg.Config.ReadArrayConfig("ACCESS_TOKEN"),
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

	s := auth2.NewService(
		hs.logger,
		hs.Secret,
		hs.JwtExpiration,
		hs.AccessSecret,
		hs.AccessContext,
		hs.AccessToken,
	)
	s.Execute(
		auth2.NewRequest(domain),
	)

	response, err := s.GetResponse()
	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(response.Status, response)
}
