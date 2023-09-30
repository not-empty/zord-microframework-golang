package http

import (
	"fmt"
	"go-skeleton/cmd/handlers/http/dependencies"
	"go-skeleton/cmd/handlers/http/routes"
	"go-skeleton/cmd/handlers/types"
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
	c := pkg.ServerDependencies["config"]
	l := pkg.ServerDependencies["logger"]

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

	deps := types.NewDependencies()
	middlewareList := dependencies.GetAllMiddlewares(hs.config.ReadConfig("JWT_SECRET"))

	for index, mid := range middlewareList {
		mid.Boot()
		pkg.Logger.Info(fmt.Sprintf("[http.Server] Booting %s", index))
	}

	protectedRoutes := routes.GetProtectedRoutes(deps)
	publicRoutes := routes.GetPublicRoutes(deps)

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
	hs.Shutdown(server.Start(port))
}

func (hs *Server) Shutdown(err error) {
	hs.logger.Critical(err, "Unable to start server, Shutdown")
}
