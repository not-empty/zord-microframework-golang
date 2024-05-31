package cmd

import (
	dummyRepository "go-skeleton/internal/repositories/dummy"
	//{{codeGen5}}
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
	"go-skeleton/pkg/validator"
)

var (
	Reg       *registry.Registry
	ApiPrefix string
)

func Setup() {
	conf := config.NewConfig()
	err := conf.LoadEnvs()
	if err != nil {
		panic(err)
	}

	ApiPrefix = conf.ReadConfig("API_PREFIX")

	l := logger.NewLogger(
		conf.ReadConfig("ENVIRONMENT"),
		conf.ReadConfig("APP"),
		conf.ReadConfig("VERSION"),
	)

	l.Boot()

	db := database.NewMysql(
		l,
		conf.ReadConfig("DB_USER"),
		conf.ReadConfig("DB_PASS"),
		conf.ReadConfig("DB_URL"),
		conf.ReadConfig("DB_PORT"),
		conf.ReadConfig("DB_DATABASE"),
	)

	idC := idCreator.NewIdCreator()
	val := validator.NewValidator()

	db.Connect()
	val.Boot()

	Reg = registry.NewRegistry()
	Reg.Provide("logger", l)
	Reg.Provide("validator", val)
	Reg.Provide("config", conf)
	Reg.Provide("idCreator", idC)

	Reg.Provide("dummyRepository", dummyRepository.NewDummyRepository(db.Db))
	//{{codeGen6}}
}
