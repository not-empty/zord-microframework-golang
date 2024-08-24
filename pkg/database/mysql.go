package database

import (
	"go-skeleton/pkg/logger"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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
	config := mysql.Config{
		User:      m.DbUser,
		Passwd:    m.DbPass,
		Addr:      m.DbUrl + ":" + m.DbPort,
		Net:       "tcp",
		ParseTime: true,
		DBName:    m.Database,
	}

	conn, err := sqlx.Open("mysql", config.FormatDSN())
	conn.SetMaxOpenConns(5)
	conn.SetMaxIdleConns(3)
	if err != nil {
		log.Fatalln(err)
	}
	m.Db = conn
}
