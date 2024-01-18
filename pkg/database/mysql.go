package database

import (
	"go-skeleton/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySql struct {
	logger   *logger.Logger
	Db       *gorm.DB
	DbConfig *DbConfig
}

func NewMysql(
	l *logger.Logger,
	DbConfig *DbConfig,
) *MySql {
	return &MySql{
		logger:   l,
		DbConfig: DbConfig,
	}
}

func (m *MySql) Connect() {
	dsn := m.DbConfig.GetDsn()
	m.DbConfig.CreateConnection(dsn)

	m.DbConfig.Connection.SetMaxOpenConns(30)
	m.DbConfig.Connection.SetMaxIdleConns(20)

	dialector := mysql.New(mysql.Config{
		Conn: m.DbConfig.Connection,
	})

	database, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		m.logger.Critical(err)
	}
	m.Db = database
}
