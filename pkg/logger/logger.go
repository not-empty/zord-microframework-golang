package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	env          string
	app          string
	version      string
	logger       zerolog.Logger
	isProduction bool
	service      string
}

func NewLogger(environment, app, version string) *Logger {
	return &Logger{
		env:          environment,
		app:          app,
		version:      version,
		isProduction: environment == "production",
	}
}

func (l *Logger) Boot() {
	zerolog.CallerFieldName = "trace"
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro

	l.logger = log.With().
		Str("app", l.app).
		Str("version", l.version).
		Str("environment", l.env).
		Logger()
}

func (l *Logger) Debug(Message string, Context ...string) {
	l.logger.Debug().Str("service", l.service).Strs("context", Context).Msg(Message)
}

func (l *Logger) Info(Message string, Context ...string) {
	l.logger.Info().
		Str("service", l.service).Strs("context", Context).Msg(Message)
}

func (l *Logger) Warning(Message string, Context ...string) {
	l.logger.Warn().
		Str("service", l.service).Strs("context", Context).Msg(Message)
}

func (l *Logger) Error(Error error, Context ...string) {
	l.logger.Error().Str("service", l.service).Strs(
		"context", Context).
		Caller(1).Msg(Error.Error())
}

func (l *Logger) Critical(Error error, Context ...string) {
	l.logger.Fatal().Str("service", l.service).Strs(
		"context", Context).
		Caller(1).Msg(Error.Error())
}

func (l *Logger) SetLogService(service string) {
	l.service = service
}
