package logger

import (
	"errors"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger("testing", "zord_app", "1.0.0")

	assert.NotNil(t, logger)
	assert.Equal(t, "testing", logger.env)
	assert.Equal(t, "zord_app", logger.app)
	assert.Equal(t, "1.0.0", logger.version)
	assert.False(t, logger.isProduction)
}

func TestBoot(t *testing.T) {
	logger := NewLogger("production", "myApp", "1.0.0")
	logger.Boot()
	assert.NotNil(t, logger.logger)
}

func TestDebug(t *testing.T) {
	logger := NewLogger("production", "myApp", "1.0.0")
	logger.Boot()
	logger.SetLogService("myService")

	logger.Debug("Debug message", "context1", "context2")

	// Act
	logger.Debug("Debug message", "context1", "context2")

	// Assert
	// 	loggedEntry := captureLastLogEntry(logger.logger)
	// 	assert.NotNil(t, loggedEntry)
	// 	assert.Equal(t, zerolog.DebugLevel, loggedEntry.Level())
	// 	assert.Contains(t, loggedEntry.Message().Str, "Debug message")
	// 	assert.Equal(t, "myService", loggedEntry.Str("service"))
	// 	assert.ElementsMatch(t, []string{"context1", "context2"}, loggedEntry.Strs("context"))
}

// Helper function to capture the last log entry
// func captureLastLogEntry(logger zerolog.Logger) *zerolog.Event {
// 	// logger.With().Logger()
// 	// logger.With().Logger().Hook()
// 	events := logger.WithContext(logger.With().Logger(nil)).Hook(&captureHook{}).With().Logger().Out.(*captureHook).entries
// 	if len(events) > 0 {
// 		return events[len(events)-1]
// 	}
// 	return nil
// }

// Helper captureHook to capture log entries
type captureHook struct {
	entries []*zerolog.Event
}

func (h *captureHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	h.entries = append(h.entries, e)
}

func TestInfo(t *testing.T) {
	logger := NewLogger("production", "myApp", "1.0.0")
	logger.Boot()
	logger.SetLogService("myService")

	logger.Info("Info message", "context1", "context2")
}

func TestWarning(t *testing.T) {
	logger := NewLogger("production", "myApp", "1.0.0")
	logger.Boot()
	logger.SetLogService("myService")

	logger.Warning("Warning message", "context1", "context2")
}

func TestError(t *testing.T) {
	logger := NewLogger("production", "myApp", "1.0.0")
	logger.Boot()
	logger.SetLogService("myService")

	err := errors.New("Some error")
	logger.Error(err, "context1", "context2")
}

func TestCritical(t *testing.T) {
	logger := NewLogger("production", "myApp", "1.0.0")
	logger.Boot()
	logger.SetLogService("myService")

	err := errors.New("Critical error")
	logger.Critical(err, "context1", "context2")
}
