package handler_test

import (
	"fmt"
	h "github.com/madeiramadeirabr/lib-go-logger/pkg/handler"
	l "github.com/madeiramadeirabr/lib-go-logger/pkg/log_level"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Run("Stdout", func(t *testing.T) {
		t.Run("Should return true when log level is handling", func(t *testing.T) {

			logLevels := [6]l.LogLevelEnum{l.LogLevelTrace, l.LogLevelDebug, l.LogLevelInfo, l.LogLevelWarning, l.LogLevelError, l.LogLevelEmergency}

			for i, logLevel := range logLevels {
				t.Run(fmt.Sprintf("Should return correct boolean when log level is %s", logLevel), func(t *testing.T) {
					handler := h.New(logLevel)
					for j, messageLevel := range logLevels {
						assert.Equal(t, j >= i, handler.IsHandling(messageLevel))
					}
				})

			}
		})

		t.Run("Should return true when write to stdout", func(t *testing.T) {

			handler := h.New(l.LogLevelDebug)
			response := handler.Write("Mensagem para o stdout")
			assert.True(t, response)
		})
	})
}
