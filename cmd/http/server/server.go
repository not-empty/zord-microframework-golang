package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-skeleton/cmd/http/middlewares"
	"go-skeleton/cmd/http/routes"
	"go-skeleton/pkg"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"
)

type Server struct {
	config *config.Config
	logger *logger.Logger
}

func NewServer() *Server {
	return &Server{
		config: pkg.Config,
		logger: pkg.Logger,
	}
}

func (hs *Server) Start() {
	var server = echo.New()

	server.HideBanner = true
	server.HidePort = true

	protectedRoutes := routes.GetProtectedRoutes()
	publicRoutes := routes.GetPublicRoutes()

	public := server.Group("")
	protected := server.Group("", middlewares.AuthMiddleware(hs.config.ReadConfig("JWT_SECRET")))

	for index, route := range publicRoutes {
		route.DeclareRoutes(public)
		hs.logger.Info(fmt.Sprintf("[server.route] Declared %s", index))
	}

	for index, route := range protectedRoutes {
		route.DeclareRoutes(protected)
		hs.logger.Info(fmt.Sprintf("[server.route] Declared %s", index))
	}
	hs.Shutdown(server.Start(":" + hs.config.ReadConfig("HTTP_PORT")))
}

func (hs *Server) Shutdown(err error) {
	hs.logger.Critical(err, "Unable to start server, Shutdown")
}
