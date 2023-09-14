package routes

import (
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/logger"

	"github.com/labstack/echo/v4"
)

type Declarable interface {
	DeclareRoutes(*echo.Echo)
}

func GetAllRoutes(logger *logger.Logger, Environment string, MySql *database.MySql) map[string]Declarable {
	dummyListRoutes := NewDummyRoutes(logger, Environment, MySql)
	//{{codeGen1}}
	domains := map[string]Declarable{
		"dummy": dummyListRoutes,
		//{{codeGen2}}
	}
	return domains
}
