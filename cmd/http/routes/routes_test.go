package routes

import (
	dummyRepository "go-skeleton/internal/repositories/dummy"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
	"go-skeleton/pkg/validator"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var conf = config.NewConfig()

var l = &logMock{}

var DbConfig = NewDbConfigMock(
	"DbUser",
	"DbPass",
	"DbUrl",
	"3306",
	"mysql",
	"Database",
	l,
)
var db, _, _ = sqlmock.New()
var mysql = database.NewMysql(l, DbConfig)

var idC = idCreator.NewIdCreator()
var val = validator.NewValidator()

func TestGetPublicRoutes(t *testing.T) {
	err := conf.LoadEnvs()
	if err != nil {
		panic(err)
	}

	l.Boot()

	DbConfig.Connection = db
	mysql.Connect()
	val.Boot()

	route := NewRoutes()
	reg := registry.NewRegistry()
	reg.Provide("logger", l)
	reg.Provide("validator", val)
	reg.Provide("config", conf)
	reg.Provide("idCreator", idC)
	reg.Provide("routes", route)

	reg.Provide("dummyRepository", dummyRepository.NewDummyRepo(mysql))

	route.GetPublicRoutes(reg)
}

type logMock struct {
}

func (l *logMock) Boot() {}

func (l *logMock) Debug(Message string, Context ...string) {}

func (l *logMock) Info(Message string, Context ...string) {}

func (l *logMock) Warning(Message string, Context ...string) {}

func (l *logMock) Error(Error error, Context ...string) {}

func (l *logMock) Critical(Error error, Context ...string) {}

func (l *logMock) SetLogService(service string) {}

type DbConfigMock struct {
	database.DbConfig
}

func NewDbConfigMock(
	DbUser string,
	DbPass string,
	DbUrl string,
	DbPort string,
	Driver string,
	Database string,
	logger logger.ILogger,
) *DbConfigMock {
	return &DbConfigMock{
		DbConfig: database.DbConfig{
			DbUser:   DbUser,
			DbPass:   DbPass,
			DbUrl:    DbUrl,
			DbPort:   DbPort,
			Driver:   Driver,
			Database: Database,
		},
	}
}

func (db *DbConfigMock) GetDsn() string {
	return "sqlmock_db_0/teste"
}

func (db *DbConfigMock) CreateConnection(dsn string) {
}

func TestDummyDeclareRoutes(t *testing.T) {
	srv := echo.New()
	group := srv.Group("")

	err := conf.LoadEnvs()
	if err != nil {
		panic(err)
	}

	l.Boot()

	DbConfig.Connection = db
	mysql.Connect()
	val.Boot()

	route := NewRoutes()
	reg := registry.NewRegistry()
	reg.Provide("logger", l)
	reg.Provide("validator", val)
	reg.Provide("config", conf)
	reg.Provide("idCreator", idC)
	reg.Provide("routes", route)

	reg.Provide("dummyRepository", dummyRepository.NewDummyRepo(mysql))

	dummyRoutes := NewDummyRoutes(reg)

	dummyRoutes.DeclareRoutes(group)
}

func TestHealthDeclareRoutes(t *testing.T) {
	srv := echo.New()
	group := srv.Group("")

	err := conf.LoadEnvs()
	if err != nil {
		panic(err)
	}

	l.Boot()

	DbConfig.Connection = db
	mysql.Connect()
	val.Boot()

	route := NewRoutes()
	reg := registry.NewRegistry()
	reg.Provide("logger", l)
	reg.Provide("validator", val)
	reg.Provide("config", conf)
	reg.Provide("idCreator", idC)
	reg.Provide("routes", route)

	dummyRoutes := NewHealthRoute()

	dummyRoutes.DeclareRoutes(group)
}

type Context struct {
	echo.Context
}

func (c *Context) JSON(code int, i interface{}) error {
	return nil
}

func TestHealth(t *testing.T) {
	ctx := &Context{}
	assert.Empty(t, health(ctx))
}
