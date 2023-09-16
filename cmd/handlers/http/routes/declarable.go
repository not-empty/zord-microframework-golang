package routes

import (
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"

	"github.com/labstack/echo/v4"
)

type Declarable interface {
	DeclareRoutes(*echo.Echo)
}

func GetAllRoutes(logger *logger.Logger, Environment string, MySql *database.MySql, idCreator *idCreator.IdCreator) map[string]Declarable {
	dummyListRoutes := NewDummyRoutes(logger, Environment, MySql, idCreator)
	//{{codeGen1}}
	domains := map[string]Declarable{
		"dummy": dummyListRoutes,
		//{{codeGen2}}
	}
	return domains
}
