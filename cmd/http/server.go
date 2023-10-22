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

	config *config.Config
	logger *logger.Logger
	deps   map[string]pkg.Bootable
}

func NewServer() *Server {
	e := pkg.Config.ReadConfig("ENVIRONMENT")
	c := pkg.ServerDependencies["config"]
	l := pkg.ServerDependencies["logger"]

	return &Server{
		Environment: e,
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

	middlewareList := middlewares.GetAllMiddlewares(hs.config.ReadConfig("JWT_SECRET"))

	for index, mid := range middlewareList {
		mid.Boot()
		pkg.Logger.Info(fmt.Sprintf("[http.Server] Booting %s", index))
	}

	protectedRoutes := routes.GetProtectedRoutes(hs.deps, hs.config.ReadConfig("ENVIRONMENT"))
	publicRoutes := routes.GetPublicRoutes(hs.deps)

	public := server.Group("")
	protected := server.Group("", middlewareList["auth"].GetMiddleware())

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
		pkg.Logger.Info(fmt.Sprintf("[Kernel.Kernel] Booting %s", index))
	}
}

func (hs *Server) Shutdown(err error) {
	hs.logger.Critical(err, "Unable to start server, Shutdown")
}
