package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-skeleton/pkg/logger"
	"log"
)

type MySql struct {
	logger   *logger.Logger
	Db       *sqlx.DB
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

func (m *MySql) Connect() {
	dsn := "%s:%s@tcp(%s:%s)/%s"
	dsn = fmt.Sprintf(
		dsn,
		m.DbUser,
		m.DbPass,
		m.DbUrl,
		m.DbPort,
		m.Database,
	)

	conn, err := sqlx.Open("mysql", dsn)
	conn.SetMaxOpenConns(30)
	conn.SetMaxIdleConns(20)
	if err != nil {
		log.Fatalln(err)
	}
	m.Db = conn
}
