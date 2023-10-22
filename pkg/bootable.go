package pkg

import (
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/validator"
)

type Bootable interface {
	Boot()
}

var Config = config.NewConfig()
var Logger = logger.NewLogger(
	Config.ReadConfig("ENVIRONMENT"),
	Config.ReadConfig("APP"),
	Config.ReadConfig("VERSION"),
)
var Mysql = database.NewMysql(
	Logger,
	Config.ReadConfig("DB_USER"),
	Config.ReadConfig("DB_PASS"),
	Config.ReadConfig("DB_URL"),
	Config.ReadConfig("DB_PORT"),
	Config.ReadConfig("DB_DATABASE"),
)

var IdCreator = idCreator.NewIdCreator()
var Validator = validator.NewValidator()

var ServerDependencies = map[string]Bootable{
	"config":    Config,
	"logger":    Logger,
	"mysql":     Mysql,
	"IdCreator": IdCreator,
	"validator": Validator,
}

var MigratorDependencies = map[string]Bootable{
	"mysql": Mysql,
}

var CliDependencies = map[string]Bootable{
	"logger": Logger,
}
