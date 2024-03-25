package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-skeleton/cmd/http/server"
	dummy2 "go-skeleton/internal/application/domain/dummy"
	dummyRepository "go-skeleton/internal/repositories/dummy"
	"log"

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

	idC := idCreator.NewIdCreator()
	val := validator.NewValidator()

	db.Connect()
	val.Boot()

	reg = registry.NewRegistry()
	reg.Provide("logger", l)
	reg.Provide("validator", val)
	reg.Provide("config", conf)
	reg.Provide("idCreator", idC)

	dsn := "%s:%s@tcp(%s:%s)/%s"
	dsn = fmt.Sprintf(
		dsn,
		conf.ReadConfig("DB_USER"),
		conf.ReadConfig("DB_PASS"),
		conf.ReadConfig("DB_URL"),
		conf.ReadConfig("DB_PORT"),
		conf.ReadConfig("DB_DATABASE"),
	)

	conn, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	conn.Exec(`
		create table dummies (
	    	id char(26) PRIMARY KEY,
	    	name char(255)
		);
	`)

	repo := dummyRepository.NewDummyRepo(conn)

	err = repo.Edit(dummy2.Dummy{
		DummyName: "Levy Sampaio",
	}, "id", "123")

	fmt.Println(err)

	reg.Provide("dummyRepository", repo)
	//{{codeGen6}}
}
