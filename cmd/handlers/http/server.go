package http

import (
	"fmt"
	"go-skeleton/cmd/handlers/http/dependencies"
	"go-skeleton/cmd/handlers/http/routes"
	"go-skeleton/cmd/handlers/types"
	"go-skeleton/pkg"

	"github.com/labstack/echo/v4"
)

func Start(port string) {
	var server = echo.New()

	hs := types.NewServer()

	server.HideBanner = true
	server.HidePort = true

	middlewareList := dependencies.GetAllMiddlewares(hs.Config.ReadConfig("JWT_SECRET"))

	for index, mid := range middlewareList {
		mid.Boot()
		pkg.Logger.Info(fmt.Sprintf("[http.Server] Booting %s", index))
	}

	protectedRoutes := routes.GetProtectedRoutes(hs)
	publicRoutes := routes.GetPublicRoutes(hs)

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

	hs.Logger.Critical(server.Start(port), "Unable to start server, Shutdown")
}
