package main

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
	Environment string
	jwtSecret   string

	config *config.Config
	logger *logger.Logger
	deps   map[string]pkg.Bootable
}

func NewServer() *Server {
	e := pkg.Config.ReadConfig("ENVIRONMENT")
	jwt := pkg.Config.ReadConfig("JWT_SECRET")
	c := pkg.ServerDependencies["config"]
	l := pkg.ServerDependencies["logger"]

	return &Server{
		Environment: e,
		jwtSecret:   jwt,
		config:      c.(*config.Config),
		logger:      l.(*logger.Logger),
		deps:        pkg.ServerDependencies,
	}
}

func main() {
	server := NewServer()
	server.Boot()
	server.Start()
}

func (hs *Server) Start() {
	var server = echo.New()

	server.HideBanner = true
	server.HidePort = true

	protectedRoutes := routes.GetProtectedRoutes(hs.deps, hs.Environment)
	publicRoutes := routes.GetPublicRoutes(hs.deps)

	public := server.Group("")
	protected := server.Group("", middlewares.AuthMiddleware(hs.jwtSecret))

	for index, route := range publicRoutes {
		route.DeclareRoutes(public)
		pkg.Logger.Info(fmt.Sprintf("[server.route] Declared %s", index))
	}

	for index, route := range protectedRoutes {
		route.DeclareRoutes(protected)
		pkg.Logger.Info(fmt.Sprintf("[server.route] Declared %s", index))
	}
	hs.Shutdown(server.Start(":" + hs.config.ReadConfig("HTTP_PORT")))
}

func (hs *Server) Boot() {
	for index, dep := range pkg.ServerDependencies {
		dep.Boot()
		pkg.Logger.Info(fmt.Sprintf("[http.server] Booting %s", index))
	}
}

func (hs *Server) Shutdown(err error) {
	hs.logger.Critical(err, "Unable to start server, Shutdown")
}
