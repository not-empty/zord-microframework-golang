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
	config   config.IConfig
	logger   logger.ILogger
	registry *registry.Registry
	routes   routes.IRoutes
	echo     IServer
}

type IServer interface {
	Start(address string) error
	Use(middleware ...echo.MiddlewareFunc)
	Group(prefix string, m ...echo.MiddlewareFunc) (g *echo.Group)
}

func NewServer(reg *registry.Registry, echo IServer) *Server {
	return &Server{
		config:   reg.Inject("config").(config.IConfig),
		logger:   reg.Inject("logger").(logger.ILogger),
		routes:   reg.Inject("routes").(routes.IRoutes),
		echo:     echo,
		registry: reg,
	}
}

func (hs *Server) Start() {
	hs.echo.Use(middleware.Recover())

	publicRoutes := hs.routes.GetPublicRoutes(hs.registry)

	public := hs.echo.Group("")

	for index, route := range publicRoutes {
		route.DeclareRoutes(public)
		hs.logger.Info(fmt.Sprintf("[server.route] Declared %s", index))
	}

	hs.Shutdown(hs.echo.Start(":" + hs.config.ReadConfig("HTTP_PORT")))
}

func (hs *Server) Shutdown(err error) {
	hs.logger.Critical(err, "Unable to start server, Shutdown")
}
