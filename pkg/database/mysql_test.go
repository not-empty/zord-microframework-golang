package database

import (
	"go-skeleton/pkg/logger"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
)

func TestNewDbConfig(t *testing.T) {
	l := logger.NewLogger("teste", "teste", "teste")
	DbConfig := NewDbConfig(
		"DbUser",
		"DbPass",
		"DbUrl",
		"DbPort",
		"mysql",
		"Database",
		l,
	)
	assert.Equal(t, "DbUser", DbConfig.DbUser)
	assert.Equal(t, "DbPass", DbConfig.DbPass)
	assert.Equal(t, "DbUrl", DbConfig.DbUrl)
	assert.Equal(t, "DbPort", DbConfig.DbPort)
	assert.Equal(t, "Database", DbConfig.Database)
}

func TestGetDsn(t *testing.T) {
	l := logger.NewLogger("teste", "teste", "teste")
	DbConfig := NewDbConfig(
		"DbUser",
		"DbPass",
		"DbUrl",
		"DbPort",
		"mysql",
		"Database",
		l,
	)

	dsn := DbConfig.GetDsn()
	assert.Equal(t, "DbUser:DbPass@tcp(DbUrl:DbPort)/Database", dsn)
}

func TestCreateConnection(t *testing.T) {
	l := logger.NewLogger("teste", "teste", "teste")
	DbConfig := NewDbConfig(
		"DbUser",
		"DbPass",
		"DbUrl",
		"DbPort",
		"mysql",
		"Database",
		l,
	)

	dsn := DbConfig.GetDsn()
	DbConfig.CreateConnection(dsn)

	assert.NotEmpty(t, DbConfig.Connection)
}

func TestCreateConnectionError(t *testing.T) {
	l := logger.NewLogger("teste", "teste", "teste")
	DbConfig := NewDbConfig(
		"DbUser",
		"DbPass",
		"DbUrl",
		"DbPort",
		"lala",
		"Database",
		l,
	)

	assert.Panics(t, func() { DbConfig.CreateConnection("teste") }, "lala")
}

func TestGetGorm(t *testing.T) {
	l := logger.NewLogger("teste", "teste", "teste")
	DbConfig := NewDbConfig(
		"DbUser",
		"DbPass",
		"DbUrl",
		"DbPort",
		"lala",
		"Database",
		l,
	)
	db, _, _ := sqlmock.New()
	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	gm := DbConfig.GetGorm(dialector)
	assert.NotEmpty(t, gm)
}

func TestGetGormError(t *testing.T) {
	l := logger.NewLogger("teste", "teste", "teste")
	DbConfig := NewDbConfig(
		"DbUser",
		"DbPass",
		"DbUrl",
		"DbPort",
		"lala",
		"Database",
		l,
	)
	dialector := mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
	})
	assert.Panics(t, func() { DbConfig.GetGorm(dialector) })
}

func TestNewMysql(t *testing.T) {
	l := logger.NewLogger("teste", "teste", "teste")
	DbConfig := NewDbConfig(
		"DbUser",
		"DbPass",
		"DbUrl",
		"DbPort",
		"mysql",
		"Database",
		l,
	)

	mysql := NewMysql(l, DbConfig)
	assert.IsType(t, &MySql{}, mysql)
}

func TestConnect(t *testing.T) {
	l := logger.NewLogger("teste", "teste", "teste")
	DbConfig := NewDbConfigMock(
		"DbUser",
		"DbPass",
		"DbUrl",
		"3306",
		"mysql",
		"Database",
		l,
	)
	db, _, _ := sqlmock.New()
	DbConfig.Connection = db
	mysql := NewMysql(l, DbConfig)
	mysql.Connect()
	assert.NotEmpty(t, mysql.Db)
}

type DbConfigMock struct {
	DbConfig
}

func NewDbConfigMock(
	DbUser string,
	DbPass string,
	DbUrl string,
	DbPort string,
	Driver string,
	Database string,
	logger *logger.Logger,
) *DbConfigMock {
	return &DbConfigMock{
		DbConfig: DbConfig{
			DbUser:   DbUser,
			DbPass:   DbPass,
			DbUrl:    DbUrl,
			DbPort:   DbPort,
			Driver:   Driver,
			Database: Database,
			logger:   logger,
		},
	}
}

func (db *DbConfigMock) GetDsn() string {
	return "sqlmock_db_0/teste"
}

func (db *DbConfigMock) CreateConnection(dsn string) {
}
