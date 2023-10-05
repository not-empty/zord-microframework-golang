package http

import (
	"fmt"
	"go-skeleton/cmd/handlers/http/dependencies"
	"go-skeleton/cmd/handlers/http/routes"
	"go-skeleton/pkg"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Dependencies *dependencies.ServerDependencies
}

func NewServer(Environment string) *Server {
	return &Server{
		Dependencies: dependencies.NewServerDependencies(),
	}
}

func (hs *Server) Start(port string) {
	var server = echo.New()

	server.HideBanner = true
	server.HidePort = true

	middlewareList := dependencies.GetAllMiddlewares(hs.Dependencies.Config.ReadConfig("JWT_SECRET"))

	for index, mid := range middlewareList {
		mid.Boot()
		pkg.Logger.Info(fmt.Sprintf("[http.Server] Booting %s", index))
	}

	protectedRoutes := routes.GetProtectedRoutes(hs.Dependencies)
	publicRoutes := routes.GetPublicRoutes(hs.Dependencies)

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
	hs.Dependencies.Logger.Critical(err, "Unable to start server, Shutdown")
}
