package database

import (
	"fmt"
	"go-skeleton/pkg/logger"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
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

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		DbConfig.CreateConnection("teste")
		defer func() {
			r := recover()
			fmt.Println(r)
			assert.NotEmpty(t, r)
		}()
	}()
	wg.Wait()
}
