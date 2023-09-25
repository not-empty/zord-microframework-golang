package routes

import (
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

func GetProtectedRoutes(logger *logger.Logger, Environment string, MySql *database.MySql, idCreator *idCreator.IdCreator, validator *validator.Validator) map[string]Declarable {
	dummyListRoutes := NewDummyRoutes(logger, Environment, MySql, idCreator, validator)
	//{{codeGen1}}
	domains := map[string]Declarable{
		"dummy": dummyListRoutes,
		//{{codeGen2}}
	}
	return domains
}

func GetPublicRoutes(logger *logger.Logger, config *config.Config) map[string]Declarable {
	health := NewHealthRoute()
	auth := NewAuthRoute(
		logger,
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
