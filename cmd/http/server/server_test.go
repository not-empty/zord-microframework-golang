package server

import (
	"errors"
	"go-skeleton/cmd/http/routes"
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	reg := registry.NewRegistry()
	l := logger.NewLogger("teste", "teste", "teste")
	conf := config.NewConfig()
	reg.Provide("logger", l)
	reg.Provide("config", conf)
	reg.Provide("routes", &rt{})
	srv := NewServer(reg, &srvMock{})
	assert.NotEmpty(t, srv)
	assert.IsType(t, &config.Config{}, srv.config)
	assert.IsType(t, &logger.Logger{}, srv.logger)
	assert.IsType(t, &registry.Registry{}, srv.registry)
}

func TestShutdown(t *testing.T) {
	reg := registry.NewRegistry()
	l := logger.NewLogger("teste", "teste", "teste")
	conf := config.NewConfig()
	reg.Provide("logger", l)
	reg.Provide("config", conf)
	reg.Provide("routes", &rt{})
	srv := NewServer(reg, &srvMock{})
	assert.Panics(t, func() { srv.Shutdown(errors.New("teste")) })
}

func TestStart(t *testing.T) {
	reg := registry.NewRegistry()
	l := &logMock{}
	conf := &ConfigMock{}
	r := &rt{}
	reg.Provide("logger", l)
	reg.Provide("config", conf)
	reg.Provide("routes", r)
	sr := &srvMock{}
	srv := NewServer(reg, sr)
	srv.Start()
}

type rt struct {
}

func (r *rt) GetPublicRoutes(reg *registry.Registry) map[string]routes.Declarable {
	return map[string]routes.Declarable{
		"teste": &DecMock{},
	}
}

type ConfigMock struct {
	config.Config
}

func (c *ConfigMock) ReadConfig(key string) string {
	return "9000"
}

type DecMock struct {
}

func (d *DecMock) DeclareRoutes(server *echo.Group) {

}

type srvMock struct {
}

func (*srvMock) Start(address string) error {
	return nil
}
func (*srvMock) Use(middleware ...echo.MiddlewareFunc) {

}
func (*srvMock) Group(prefix string, m ...echo.MiddlewareFunc) (g *echo.Group) {
	return &echo.Group{}
}

type logMock struct {
}

func (l *logMock) Boot() {}

func (l *logMock) Debug(Message string, Context ...string) {}

func (l *logMock) Info(Message string, Context ...string) {}

func (l *logMock) Warning(Message string, Context ...string) {}

func (l *logMock) Error(Error error, Context ...string) {}

func (l *logMock) Critical(Error error, Context ...string) {}

func (l *logMock) SetLogService(service string) {}
