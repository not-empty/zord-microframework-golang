package database

import (
	"fmt"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySql struct {
	logger *logger.Logger
	Db     *gorm.DB
	config *config.Config
}

func NewMysql(l *logger.Logger, c *config.Config) *MySql {
	return &MySql{
		logger: l,
		config: c,
	}
}

func (m *MySql) Boot() {
	dsn := "%s:%s@tcp(%s:%s)/%s"
	dsn = fmt.Sprintf(
		dsn,
		m.config.DbUser,
		m.config.DbPass,
		m.config.DbUrl,
		m.config.DbPort,
		m.config.Database,
	)
	dialector := mysql.New(mysql.Config{
		DSN: dsn,
	})
	database, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		m.logger.Critical(err)
	}
	m.Db = database
}
