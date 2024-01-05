package server

import (
	"fmt"
	"go-skeleton/cmd/http/routes"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"

	"github.com/labstack/echo/v4"
)

type Server struct {
	config   *config.Config
	logger   *logger.Logger
	registry *registry.Registry
}

func NewServer(reg *registry.Registry) *Server {
	return &Server{
		config:   reg.Inject("config").(*config.Config),
		logger:   reg.Inject("logger").(*logger.Logger),
		registry: reg,
	}
}

func (hs *Server) Start() {
	var server = echo.New()

	server.HideBanner = true
	server.HidePort = true

	protectedRoutes := routes.GetProtectedRoutes(hs.registry)
	publicRoutes := routes.GetPublicRoutes(hs.registry)

	// middlewares.AuthMiddleware(hs.config.ReadConfig("JWT_SECRET"))
	public := server.Group("")
	protected := server.Group("")

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
