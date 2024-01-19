package database

import (
	"go-skeleton/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySql struct {
	logger   logger.ILogger
	Db       *gorm.DB
	DbConfig IDbConfig
}

func NewMysql(
	l logger.ILogger,
	DbConfig IDbConfig,
) *MySql {
	return &MySql{
		logger:   l,
		DbConfig: DbConfig,
	}
}

func (m *MySql) Connect() {
	dsn := m.DbConfig.GetDsn()
	m.DbConfig.CreateConnection(dsn)
	conn := m.DbConfig.GetConnection()
	conn.SetMaxOpenConns(30)
	conn.SetMaxIdleConns(20)

	dialector := mysql.New(mysql.Config{
		Conn:                      conn,
		SkipInitializeWithVersion: true,
	})

	database := m.DbConfig.GetGorm(dialector)
	m.Db = database
}
