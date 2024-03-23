package main

import (
	"context"
	"fmt"
	"go-skeleton/cmd/http/server"
	dummyRepository "go-skeleton/internal/repositories/dummy"
	"go-skeleton/internal/repositories/user"
	"go-skeleton/pkg/ent"

	//{{codeGen5}}
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/database"
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

	client, err := ent.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			conf.ReadConfig("DB_USER"),
			conf.ReadConfig("DB_PASS"),
			conf.ReadConfig("DB_URL"),
			conf.ReadConfig("DB_PORT"),
			conf.ReadConfig("DB_DATABASE"),
		),
	)

	errMigration := client.Schema.Create(context.Background())
	if errMigration != nil {
		return
	}

	userRepo := user.NewUserRepository(client)

	idC := idCreator.NewIdCreator()
	val := validator.NewValidator()

	db.Connect()
	val.Boot()

	reg = registry.NewRegistry()
	reg.Provide("logger", l)
	reg.Provide("validator", val)
	reg.Provide("config", conf)
	reg.Provide("idCreator", idC)

	reg.Provide("dummyRepository", dummyRepository.NewDummyRepo(db))
	reg.Provide("userRepo", userRepo)
	//{{codeGen6}}
}
