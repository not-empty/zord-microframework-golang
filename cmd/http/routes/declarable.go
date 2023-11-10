package routes

import (
	"go-skeleton/pkg/registry"

	"github.com/labstack/echo/v4"
)

type Declarable interface {
	DeclareRoutes(*echo.Group)
}

func GetProtectedRoutes(reg *registry.Registry) map[string]Declarable {
	dummyListRoutes := NewDummyRoutes(reg)
	//{{codeGen1}}
	domains := map[string]Declarable{
		"dummy": dummyListRoutes,
		//{{codeGen2}}
	}
	return domains
}

func GetPublicRoutes(reg *registry.Registry) map[string]Declarable {
	health := NewHealthRoute()
	auth := NewAuthRoute(reg)
	routes := map[string]Declarable{
		"health": health,
		"auth":   auth,
	}
	return routes
}
