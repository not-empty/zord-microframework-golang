package routes

import (
	"encoding/json"
	"go-skeleton/application/domain/auth"
	login "go-skeleton/application/services/auth/LOGIN"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"
	"io"

	"github.com/labstack/echo/v4"
)

type Auth struct {
	logger *logger.Logger
	config *config.Config
}

func NewAuthRoute(logger *logger.Logger, config *config.Config) *Auth {
	return &Auth{
		logger: logger,
		config: config,
	}
}

func (hs *Auth) DeclareRoutes(server *echo.Group) {
	server.POST("/login", hs.handleLogin)
}

func (hs *Auth) handleLogin(c echo.Context) error {
	auth := auth.Token{}

	body, errors := io.ReadAll(c.Request().Body)
	if errors != nil {
		return c.JSON(422, errors)
	}

	errors = json.Unmarshal(body, &auth)
	if errors != nil {
		return c.JSON(422, errors)
	}

	s := login.NewService(hs.logger, hs.config)
	s.Execute(
		login.NewRequest(auth),
	)

	response, err := s.GetResponse()
	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(response.Status, response)
}
