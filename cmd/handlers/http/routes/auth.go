package routes

import (
	"encoding/json"
	"go-skeleton/application/domain/auth"
	login "go-skeleton/application/services/auth/LOGIN"
	"go-skeleton/pkg/logger"
	"io"

	"github.com/labstack/echo/v4"
)

type Auth struct {
	logger        *logger.Logger
	Secret        string
	JwtExpiration int
	AccessSecret  []string
	AccessContext []string
	AccessToken   []string
}

func NewAuthRoute(
	logger *logger.Logger,
	Secret string,
	JwtExpiration int,
	AccessSecret []string,
	AccessContext []string,
	AccessToken []string,
) *Auth {
	return &Auth{
		logger:        logger,
		Secret:        Secret,
		JwtExpiration: JwtExpiration,
		AccessSecret:  AccessSecret,
		AccessContext: AccessContext,
		AccessToken:   AccessToken,
	}
}

func (hs *Auth) DeclareRoutes(server *echo.Group) {
	server.POST("/login", hs.handleLogin)
}

func (hs *Auth) handleLogin(c echo.Context) error {
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
