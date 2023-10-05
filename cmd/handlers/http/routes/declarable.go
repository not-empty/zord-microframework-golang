package routes

import (
	"github.com/labstack/echo/v4"
	"go-skeleton/pkg"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/validator"
)

type Declarable interface {
	DeclareRoutes(*echo.Group)
}

func GetProtectedRoutes(deps map[string]pkg.Bootable, Env string) map[string]Declarable {
	logger := deps["logger"].(*logger.Logger)
	mysql := deps["mysql"].(*database.MySql)
	idCreator := deps["idCreator"].(*idCreator.IdCreator)
	validator := deps["validator"].(*validator.Validator)

	dummyListRoutes := NewDummyRoutes(
		logger,
		mysql,
		idCreator,
		validator,
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
	config := deps["config"].(*config.Config)
	health := NewHealthRoute()
	auth := NewAuthRoute(
		deps["logger"].(*logger.Logger),
		config.ReadConfig("JWT_SECRET"),
		config.ReadNumberConfig("JWT_EXPIRATION"),
		config.ReadArrayConfig("ACCESS_SECRET"),
		config.ReadArrayConfig("ACCESS_CONTEXT"),
		config.ReadArrayConfig("ACCESS_TOKEN"),
	)
	routes := map[string]Declarable{
		"health": health,
		"auth":   auth,
	}
	return routes
}
