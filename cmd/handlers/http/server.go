package http

import (
	"fmt"
	"go-skeleton/cmd/handlers/http/routes"
	"go-skeleton/pkg"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/validator"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Environment string

	config    *config.Config
	logger    *logger.Logger
	mysql     *database.MySql
	idCreator *idCreator.IdCreator
	validator *validator.Validator
}

func NewServer(Environment string) *Server {
	c := pkg.ServerDependencies["config"]
	l := pkg.ServerDependencies["logger"]
	m := pkg.ServerDependencies["mysql"]
	i := pkg.ServerDependencies["IdCreator"]
	v := pkg.ServerDependencies["validator"]

	return &Server{
		Environment: Environment,
		config:      c.(*config.Config),
		logger:      l.(*logger.Logger),
		mysql:       m.(*database.MySql),
		idCreator:   i.(*idCreator.IdCreator),
		validator:   v.(*validator.Validator),
	}
}

func (hs *Server) Start(port string) {
	var server = echo.New()

	server.HideBanner = true
	server.HidePort = true
	protectedRoutes := routes.GetProtectedRoutes(hs.logger, hs.Environment, hs.mysql, hs.idCreator, hs.validator)
	publicRoutes := routes.GetPublicRoutes()

	public := server.Group("")
	protected := server.Group("/v1")

	for index, route := range publicRoutes {
		route.DeclareRoutes(public)
		pkg.Logger.Info(fmt.Sprintf("[server.route] Declared %s", index))
	}

	for index, route := range protectedRoutes {
		route.DeclareRoutes(protected)
		pkg.Logger.Info(fmt.Sprintf("[server.route] Declared %s", index))
	}
	hs.Shutdown(server.Start(port))

}

func (hs *Server) Shutdown(err error) {
	hs.logger.Critical(err, "Unable to start server, Shutdown")
}
