package main

import (
	"go-skeleton/cmd/http/server"
	dummyRepository "go-skeleton/internal/repositories/dummy"
	"go-skeleton/pkg/database"
	//{{codeGen5}}
	_ "github.com/go-sql-driver/mysql"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
	"go-skeleton/pkg/validator"
)

var (
	reg *registry.Registry
)

func main() {
	serverInstance := server.NewServer(reg)
	serverInstance.Start()
}

func init() {
	conf := config.NewConfig()
	err := conf.LoadEnvs()
	if err != nil {
		panic(err)
	}

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

	reg = registry.NewRegistry()
	reg.Provide("logger", l)
	reg.Provide("validator", val)
	reg.Provide("config", conf)
	reg.Provide("idCreator", idC)

	repo := dummyRepository.NewDummyRepo(db.Db)

	reg.Provide("dummyRepository", repo)
	//{{codeGen6}}
}
