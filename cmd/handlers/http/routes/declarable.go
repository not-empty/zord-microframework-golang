package routes

import (
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/validator"

	"github.com/labstack/echo/v4"
)

type Declarable interface {
	DeclareRoutes(*echo.Echo)
}

func GetAllRoutes(logger *logger.Logger, Environment string, MySql *database.MySql, idCreator *idCreator.IdCreator, validator *validator.Validator) map[string]Declarable {
	health := NewHealthRoute()
	dummyListRoutes := NewDummyRoutes(logger, Environment, MySql, idCreator, validator)
	//{{codeGen1}}
	domains := map[string]Declarable{
		"health": health,
		"dummy":  dummyListRoutes,
		//{{codeGen2}}
	}
	return domains
}
