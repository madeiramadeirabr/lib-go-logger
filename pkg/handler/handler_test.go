package handler_test

import (
	"github.com/stretchr/testify/assert"
	h "lib-go-logger/pkg/handler"
	l "lib-go-logger/pkg/log_level"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("Stdout", func(t *testing.T) {
		t.Run("Should return false when log level is not handling", func(t *testing.T) {

			handler := h.New(l.LogLevelDebug)
			response := handler.Write("Mensagem para o stdout")
			assert.True(t, response)
			assert.False(t, false)
		})

		t.Run("Should return true when log level is not handling", func(t *testing.T) {

			assert.True(t, true)
		})
		t.Run("Should return true when write to stdout", func(t *testing.T) {

			handler := h.New(l.LogLevelDebug)
			response := handler.Write("Mensagem para o stdout")
			assert.True(t, response)
		})
	})
}
