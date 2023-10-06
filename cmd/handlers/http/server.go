package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"go-skeleton/cmd/handlers/http/middlewares"
	"go-skeleton/cmd/handlers/http/routes"
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

func (hs *Server) Boot(_ *cobra.Command, _ []string) {
	for index, dep := range pkg.ServerDependencies {
		dep.Boot()
		pkg.Logger.Info(fmt.Sprintf("[http.Server] Booting %s", index))
	}

	pkg.Logger.Info("[http.Server] Done!")
}

func (hs *Server) Start(_ *cobra.Command, _ []string) {
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
	hs.Shutdown(server.Start(hs.config.ReadConfig("HTTP_PORT")))
}

func (hs *Server) Shutdown(err error) {
	hs.logger.Critical(err, "Unable to start server, Shutdown")
}

func (hs *Server) BaseCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "http",
		Short:             "Start a http server (API)",
		Long:              ``,
		Run:               hs.Boot,
		PersistentPostRun: hs.Start,
	}
}
