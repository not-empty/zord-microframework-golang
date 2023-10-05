package routes

import (
	"go-skeleton/cmd/handlers/http/dependencies"

	"github.com/labstack/echo/v4"
)

type Declarable interface {
	DeclareRoutes(*echo.Group)
}

func GetProtectedRoutes(deps *dependencies.ServerDependencies) map[string]Declarable {
	dummyListRoutes := NewDummyRoutes(deps)
	//{{codeGen1}}
	domains := map[string]Declarable{
		"dummy": dummyListRoutes,
		//{{codeGen2}}
	}

	return domains
}

func GetPublicRoutes(deps *dependencies.ServerDependencies) map[string]Declarable {
	health := NewHealthRoute()
	auth := NewAuthRoute(
		deps.Logger,
		deps.Config.ReadConfig("JWT_SECRET"),
		deps.Config.ReadNumberConfig("JWT_EXPIRATION"),
		deps.Config.ReadArrayConfig("ACCESS_SECRET"),
		deps.Config.ReadArrayConfig("ACCESS_CONTEXT"),
		deps.Config.ReadArrayConfig("ACCESS_TOKEN"),
	)

	routes := map[string]Declarable{
		"health": health,
		"auth":   auth,
	}

	return routes
}
