package routes

import (
	"go-skeleton/pkg"

	"github.com/labstack/echo/v4"
)

type Declarable interface {
	DeclareRoutes(*echo.Group)
}

func GetProtectedRoutes(deps map[string]pkg.Bootable, Env string) map[string]Declarable {

	dummyListRoutes := NewDummyRoutes(Env, deps)
	//{{codeGen1}}
	domains := map[string]Declarable{
		"dummy": dummyListRoutes,
		//{{codeGen2}}
	}
	return domains
}

func GetPublicRoutes(deps map[string]pkg.Bootable) map[string]Declarable {
	health := NewHealthRoute()
	auth := NewAuthRoute(deps)
	routes := map[string]Declarable{
		"health": health,
		"auth":   auth,
	}
	return routes
}
