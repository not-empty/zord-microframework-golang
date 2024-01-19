package database

import (
	"database/sql"
	"fmt"
	"go-skeleton/pkg/logger"

	"gorm.io/gorm"
)

type DbConfig struct {
	DbUser     string
	DbPass     string
	DbUrl      string
	DbPort     string
	Driver     string
	Database   string
	Connection *sql.DB
	Dialector  gorm.Dialector
	logger     logger.ILogger
}

type IDbConfig interface {
	CreateConnection(dsn string)
	GetDsn() string
	GetConnection() *sql.DB
	GetGorm(dialector gorm.Dialector) *gorm.DB
}

func NewDbConfig(
	DbUser string,
	DbPass string,
	DbUrl string,
	DbPort string,
	Driver string,
	Database string,
	logger logger.ILogger,
) *DbConfig {
	return &DbConfig{
		DbUser:   DbUser,
		DbPass:   DbPass,
		DbUrl:    DbUrl,
		DbPort:   DbPort,
		Driver:   Driver,
		Database: Database,
		logger:   logger,
	}
}

func (db *DbConfig) CreateConnection(dsn string) {
	conn, err := sql.Open(db.Driver, dsn)
	if err != nil {
		db.logger.Critical(err)
	}
	db.Connection = conn
}

func (db *DbConfig) GetDsn() string {
	dsn := "%s:%s@tcp(%s:%s)/%s"
	return fmt.Sprintf(
		dsn,
		db.DbUser,
		db.DbPass,
		db.DbUrl,
		db.DbPort,
		db.Database,
	)
}

func (db *DbConfig) GetConnection() *sql.DB {
	return db.Connection
}

func (db *DbConfig) GetGorm(dialector gorm.Dialector) *gorm.DB {
	database, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		db.logger.Critical(err)
	}
	return database
}
