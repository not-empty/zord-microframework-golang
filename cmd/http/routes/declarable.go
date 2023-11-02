package routes

import (
	"github.com/labstack/echo/v4"
)

type Declarable interface {
	DeclareRoutes(*echo.Group)
}

func GetProtectedRoutes() map[string]Declarable {

	dummyListRoutes := NewDummyRoutes()
	testeListRoutes := NewTesteRoutes(Env, deps)
	//{{codeGen1}}
	domains := map[string]Declarable{
		"dummy": dummyListRoutes,
		"teste": testeListRoutes,
		//{{codeGen2}}
	}
	return domains
}

func GetPublicRoutes() map[string]Declarable {
	health := NewHealthRoute()
	auth := NewAuthRoute()
	routes := map[string]Declarable{
		"health": health,
		"auth":   auth,
	}
	return routes
}
