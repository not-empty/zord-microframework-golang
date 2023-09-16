package http

import (
	"fmt"
	"go-skeleton/cmd/handlers/http/routes"
	"go-skeleton/pkg"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Environment string

	config    *config.Config
	logger    *logger.Logger
	mysql     *database.MySql
	idCreator *idCreator.IdCreator
}

func NewServer(Environment string) *Server {
	c := pkg.ServerDependencies["config"]
	l := pkg.ServerDependencies["logger"]
	m := pkg.ServerDependencies["mysql"]
	i := pkg.ServerDependencies["IdCreator"]

	return &Server{
		Environment: Environment,
		config:      c.(*config.Config),
		logger:      l.(*logger.Logger),
		mysql:       m.(*database.MySql),
		idCreator:   i.(*idCreator.IdCreator),
	}
}

func (hs *Server) Start(port string) {
	var server = echo.New()

	server.HideBanner = true
	server.HidePort = true
	routes := routes.GetAllRoutes(hs.logger, hs.Environment, hs.mysql, hs.idCreator)

	for index, route := range routes {
		route.DeclareRoutes(server)
		pkg.Logger.Info(fmt.Sprintf("[server.route] Declared %s", index))
	}
	hs.Shutdown(server.Start(port))

}

func (hs *Server) Shutdown(err error) {
	hs.logger.Critical(err, "Unable to start server, Shutdown")
}
