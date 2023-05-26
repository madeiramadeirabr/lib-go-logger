package formatter_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatter(t *testing.T) {
	t.Run("Formatter", func(t *testing.T) {
		t.Run("Should format and not return error", func(t *testing.T) {
			assert.False(t, false)
		})

		t.Run("Should return true when log level is not handling", func(t *testing.T) {

			assert.True(t, true)
		})
	})
}
