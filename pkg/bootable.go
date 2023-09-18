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
	Config.Environment,
	Config.App,
	Config.Version,
)
var Mysql = database.NewMysql(Logger, Config)
var IdCreator = idCreator.NewIdCreator()
var Validator = validator.NewValidator()

var KernelDependencies = map[string]Bootable{
	"config": Config,
	"logger": Logger,
}

var ServerDependencies = map[string]Bootable{
	"config":    Config,
	"logger":    Logger,
	"mysql":     Mysql,
	"IdCreator": IdCreator,
	"validator": Validator,
}

var CliDependencies = map[string]Bootable{
	"config": Config,
	"logger": Logger,
	"mysql":  Mysql,
}
