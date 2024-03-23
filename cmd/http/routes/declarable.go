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
	userRoute := NewUserRoutes(reg)
	cartRoute := NewCartRoutes(reg)
	//{{codeGen1}}
	routes := map[string]Declarable{
		"health": health,
		"dummy":  dummyListRoutes,
		"user":   userRoute,
		"cart":   cartRoute,
		//{{codeGen2}}
	}
	return routes
}
