package routes

import (
	"go-skeleton/pkg/registry"

	"github.com/labstack/echo/v4"
)

type Declarable interface {
	DeclareRoutes(*echo.Group)
}

func GetPublicRoutes(reg *registry.Registry) map[string]Declarable {
	health := NewHealthRoute()
	dummyListRoutes := NewDummyRoutes(reg)
	//{{codeGen1}}
	routes := map[string]Declarable{
		"health": health,
		"dummy":  dummyListRoutes,
		//{{codeGen2}}
	}
	return routes
}
