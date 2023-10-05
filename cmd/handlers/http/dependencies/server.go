package dependencies

import (
	"go-skeleton/pkg"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/validator"
)

type ServerDependencies struct {
	Environment string
	Config      *config.Config
	Logger      *logger.Logger
	Mysql       *database.MySql
	IdCreator   *idCreator.IdCreator
	Validator   *validator.Validator
}

func NewServerDependencies() *ServerDependencies {
	c := pkg.ServerDependencies["config"]
	l := pkg.ServerDependencies["logger"]
	m := pkg.ServerDependencies["mysql"]
	i := pkg.ServerDependencies["IdCreator"]
	v := pkg.ServerDependencies["validator"]

	return &ServerDependencies{
		Environment: pkg.Config.ReadConfig("ENVIRONMENT"),
		Config:      c.(*config.Config),
		Logger:      l.(*logger.Logger),
		Mysql:       m.(*database.MySql),
		IdCreator:   i.(*idCreator.IdCreator),
		Validator:   v.(*validator.Validator),
	}
}
