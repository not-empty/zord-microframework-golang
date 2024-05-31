package server

import (
	"fmt"
	"go-skeleton/cmd/http/routes"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config    *config.Config
	logger    *logger.Logger
	registry  *registry.Registry
	apiPrefix string
}

func NewServer(reg *registry.Registry, apiPrefix string) *Server {
	return &Server{
		config:    reg.Inject("config").(*config.Config),
		logger:    reg.Inject("logger").(*logger.Logger),
		registry:  reg,
		apiPrefix: apiPrefix,
	}
}

func (hs *Server) Start() {
	var server = echo.New()

	server.HideBanner = true
	server.HidePort = true

	server.Use(middleware.Recover())

	public := server.Group("")
	private := server.Group("")

	allRoutes := routes.GetRoutes(hs.registry)

	for index, route := range allRoutes {
		route.DeclarePublicRoutes(public, hs.apiPrefix)
		route.DeclarePrivateRoutes(private, hs.apiPrefix)
		hs.logger.Info(fmt.Sprintf("[server.route] Declared %s", index))
	}

	hs.Shutdown(server.Start(":" + hs.config.ReadConfig("HTTP_PORT")))
}

func (hs *Server) Shutdown(err error) {
	hs.logger.Critical(err, "Unable to start server, Shutdown")
}
