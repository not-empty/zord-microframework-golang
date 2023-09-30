package routes

import (
	"go-skeleton/cmd/handlers/types"

	"github.com/labstack/echo/v4"
)

type Declarable interface {
	DeclareRoutes(*echo.Group)
}

func GetProtectedRoutes(server *types.Server) map[string]Declarable {
	dummyListRoutes := NewDummyRoutes(server)
	//{{codeGen1}}
	domains := map[string]Declarable{
		"dummy": dummyListRoutes,
		//{{codeGen2}}
	}

	return domains
}

func GetPublicRoutes(server *types.Server) map[string]Declarable {
	health := NewHealthRoute()
	auth := NewAuthRoute(
		server.Logger,
		server.Config.ReadConfig("JWT_SECRET"),
		server.Config.ReadNumberConfig("JWT_EXPIRATION"),
		server.Config.ReadArrayConfig("ACCESS_SECRET"),
		server.Config.ReadArrayConfig("ACCESS_CONTEXT"),
		server.Config.ReadArrayConfig("ACCESS_TOKEN"),
	)

	routes := map[string]Declarable{
		"health": health,
		"auth":   auth,
	}

	return routes
}
