package log_level_test

import (
	"fmt"
	"testing"

	l "github.com/madeiramadeirabr/lib-go-logger/log_level"
	"github.com/stretchr/testify/assert"
)

func TestLogLevel(t *testing.T) {
	t.Run("LogLevel", func(t *testing.T) {

		t.Run("Should convert log level enum to int corretly", func(t *testing.T) {

			logLevels := [6]l.LogLevelEnum{l.LogLevelTrace, l.LogLevelDebug, l.LogLevelInfo, l.LogLevelWarning, l.LogLevelError, l.LogLevelEmergency}

			for i, logLevel := range logLevels {
				t.Run(fmt.Sprintf("Should convert %s to %d", logLevel, i), func(t *testing.T) {
					assert.Equal(t, i, l.LevelStringToInt(logLevel))
				})
			}
		})
	})
}
