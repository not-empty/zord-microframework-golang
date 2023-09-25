package database

import (
	"database/sql"
	"fmt"
	"go-skeleton/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySql struct {
	logger   *logger.Logger
	Db       *gorm.DB
	DbUser   string
	DbPass   string
	DbUrl    string
	DbPort   string
	Database string
}

func NewMysql(
	l *logger.Logger,
	DbUser string,
	DbPass string,
	DbUrl string,
	DbPort string,
	Database string,
) *MySql {
	return &MySql{
		logger:   l,
		DbUser:   DbUser,
		DbPass:   DbPass,
		DbUrl:    DbUrl,
		DbPort:   DbPort,
		Database: Database,
	}
}

func (m *MySql) Boot() {
	dsn := "%s:%s@tcp(%s:%s)/%s"
	dsn = fmt.Sprintf(
		dsn,
		m.DbUser,
		m.DbPass,
		m.DbUrl,
		m.DbPort,
		m.Database,
	)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		m.logger.Critical(err)
	}
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetMaxIdleConns(20)
	dialector := mysql.New(mysql.Config{
		Conn: sqlDB,
	})
	database, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		m.logger.Critical(err)
	}
	m.Db = database
}
