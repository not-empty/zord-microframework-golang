package routes

import (
	"go-skeleton/pkg"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/validator"

	"github.com/labstack/echo/v4"
)

type Declarable interface {
	DeclareRoutes(*echo.Group)
}

func GetProtectedRoutes(deps map[string]pkg.Bootable, Env string) map[string]Declarable {
	l := deps["logger"].(*logger.Logger)
	m := deps["mysql"].(*database.MySql)
	i := deps["IdCreator"].(*idCreator.IdCreator)
	v := deps["validator"].(*validator.Validator)

	dummyListRoutes := NewDummyRoutes(
		l,
		m,
		i,
		v,
		Env,
	)
	//{{codeGen1}}
	domains := map[string]Declarable{
		"dummy": dummyListRoutes,
		//{{codeGen2}}
	}
	return domains
}

func GetPublicRoutes(deps map[string]pkg.Bootable) map[string]Declarable {
	c := deps["config"].(*config.Config)
	health := NewHealthRoute()
	auth := NewAuthRoute(
		deps["logger"].(*logger.Logger),
		c.ReadConfig("JWT_SECRET"),
		c.ReadNumberConfig("JWT_EXPIRATION"),
		c.ReadArrayConfig("ACCESS_SECRET"),
		c.ReadArrayConfig("ACCESS_CONTEXT"),
		c.ReadArrayConfig("ACCESS_TOKEN"),
	)
	routes := map[string]Declarable{
		"health": health,
		"auth":   auth,
	}
	return routes
}
