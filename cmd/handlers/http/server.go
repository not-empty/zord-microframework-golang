package http

import (
	"fmt"
	"go-skeleton/cmd/handlers/http/routes"
	"go-skeleton/pkg"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Environment string

	config *config.Config
	logger *logger.Logger
}

func NewServer(Environment string) *Server {
	c := pkg.KernelDependencies["config"]
	l := pkg.KernelDependencies["logger"]

	return &Server{
		Environment: Environment,
		config:      c.(*config.Config),
		logger:      l.(*logger.Logger),
	}
}

func (hs *Server) Start(port string) {
	var server = echo.New()

	server.HideBanner = true
	server.HidePort = true
	routes := routes.GetAllRoutes(hs.logger, hs.Environment)

	for index, route := range routes {
		route.DeclareRoutes(server)
		pkg.Logger.Info(fmt.Sprintf("[server.route] Declared %s", index))
	}
	hs.Shutdown(server.Start(port))

}

func (hs *Server) Shutdown(err error) {
	hs.logger.Critical(err, "Unable to start server, Shutdown")
}
