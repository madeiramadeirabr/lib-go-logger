package logger_test

import (
	"testing"

	"errors"
	"github.com/stretchr/testify/assert"
	f "github.com/madeiramadeirabr/lib-go-logger/pkg/formatter"
	h "github.com/madeiramadeirabr/lib-go-logger/pkg/handler"
	level "github.com/madeiramadeirabr/lib-go-logger/pkg/log_level"
	l "github.com/madeiramadeirabr/lib-go-logger/pkg/logger"
)

func TestLogger(t *testing.T) {
	t.Run("Log", func(t *testing.T) {
		t.Run("Should return false when log level is not handling", func(t *testing.T) {
			handlerMock := new(h.Mock)
			handlerMock.On("IsHandling", level.LogLevelEmergency).Return(false)
			formatterMock := new(f.Mock)

			logger := l.New(
				handlerMock,
				formatterMock,
			)

			response := logger.Log(level.LogLevelEmergency, "mensagem", f.LogMessageOptions{})

			defer handlerMock.ClearExpectation()
			assert.False(t, response)
		})

		t.Run("Should return false when log level is handling but format returns error", func(t *testing.T) {
			handlerMock := new(h.Mock)
			handlerMock.On("IsHandling", level.LogLevelEmergency).Return(true)
			formatterMock := new(f.Mock)
			formatterMock.On("Format", level.LogLevelEmergency, "mensagem", f.LogMessageOptions{}).Return("", errors.New("Erro qualquer"))

			logger := l.New(
				handlerMock,
				formatterMock,
			)

			response := logger.Log(level.LogLevelEmergency, "mensagem", f.LogMessageOptions{})

			defer handlerMock.ClearExpectation()
			defer formatterMock.ClearExpectation()
			assert.False(t, response)
		})

		t.Run("Should return true when log level is handling and formatter success format", func(t *testing.T) {
			handlerMock := new(h.Mock)
			handlerMock.On("IsHandling", level.LogLevelEmergency).Return(true)
			handlerMock.On("Write", "mensagem").Return(true)
			formatterMock := new(f.Mock)
			formatterMock.On("Format", level.LogLevelEmergency, "mensagem", f.LogMessageOptions{}).Return("mensagem", nil)

			logger := l.New(
				handlerMock,
				formatterMock,
			)

			response := logger.Log(level.LogLevelEmergency, "mensagem", f.LogMessageOptions{})

			defer handlerMock.ClearExpectation()
			defer formatterMock.ClearExpectation()
			assert.True(t, response)
		})

		t.Run("Should call IsHandling, Format and Write when all functions success", func(t *testing.T) {
			handlerMock := new(h.Mock)
			handlerMock.On("IsHandling", level.LogLevelEmergency).Return(true)
			handlerMock.On("Write", "mensagem").Return(true)

			formatterMock := new(f.Mock)
			formatterMock.On("Format", level.LogLevelEmergency, "mensagem", f.LogMessageOptions{}).Return("mensagem", nil)

			logger := l.New(
				handlerMock,
				formatterMock,
			)

			_ = logger.Log(level.LogLevelEmergency, "mensagem", f.LogMessageOptions{})
			defer handlerMock.ClearExpectation()
			defer formatterMock.ClearExpectation()
			handlerMock.AssertCalled(t, "IsHandling", level.LogLevelEmergency)
			handlerMock.AssertCalled(t, "Write", "mensagem")
			formatterMock.AssertCalled(t, "Format", level.LogLevelEmergency, "mensagem", f.LogMessageOptions{})
		})
	})
}
