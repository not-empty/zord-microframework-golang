package database

import (
	"database/sql"
	"fmt"
	"go-skeleton/pkg/logger"
)

type DbConfig struct {
	DbUser     string
	DbPass     string
	DbUrl      string
	DbPort     string
	Driver     string
	Database   string
	Connection *sql.DB
	logger     *logger.Logger
}

func NewDbConfig(
	DbUser string,
	DbPass string,
	DbUrl string,
	DbPort string,
	Driver string,
	Database string,
	logger *logger.Logger,
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
	db.Connection = conn
	if err != nil {
		db.logger.Critical(err)
	}
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
