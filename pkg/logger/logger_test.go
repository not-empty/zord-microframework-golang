package logger

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type capturedLogs struct {
	Logs []string
	io.Writer
}

func (c *capturedLogs) Write(msg []byte) (int, error) {
	c.Logs = append(c.Logs, string(msg))
	return len(msg), nil
}

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

	logs := &capturedLogs{}
	logger.logger = logger.logger.Output(logs)
	logger.Debug("Debug message", "context1", "context2")
	expectedLog := `"message":"Debug message"`

	assert.Contains(t, logs.Logs[0], expectedLog)
}

func TestInfo(t *testing.T) {
	logger := NewLogger("production", "myApp", "1.0.0")
	logger.Boot()
	logger.SetLogService("myService")

	logs := &capturedLogs{}
	logger.logger = logger.logger.Output(logs)
	logger.Info("Info message", "context1", "context2")
	expectedLog := `"message":"Info message"`
	assert.Contains(t, logs.Logs[0], expectedLog)
}

func TestWarning(t *testing.T) {
	logger := NewLogger("production", "myApp", "1.0.0")
	logger.Boot()
	logger.SetLogService("myService")

	logs := &capturedLogs{}
	logger.logger = logger.logger.Output(logs)
	logger.Warning("Warning message", "context1", "context2")

	expectedLog := `"message":"Warning message"`
	assert.Contains(t, logs.Logs[0], expectedLog)
}

func TestError(t *testing.T) {
	logger := NewLogger("production", "myApp", "1.0.0")
	logger.Boot()
	logger.SetLogService("myService")

	logs := &capturedLogs{}
	logger.logger = logger.logger.Output(logs)
	err := errors.New("Some error")
	logger.Error(err, "context1", "context2")

	expectedLog := `"message":"Some error"`
	assert.Contains(t, logs.Logs[0], expectedLog)
}

func TestCritical(t *testing.T) {
	logger := NewLogger("production", "myApp", "1.0.0")
	logger.Boot()
	logger.SetLogService("myService")

	err := errors.New("Critical error")

	go func() {
		defer func() {
			r := recover()
			assert.Contains(t, r, "Critical error")
		}()
		logger.Critical(err, "context1", "context2")
	}()
}
