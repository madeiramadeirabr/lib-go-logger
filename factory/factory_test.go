package factory_test

import (
	"github.com/madeiramadeirabr/lib-go-logger/factory"
	"github.com/madeiramadeirabr/lib-go-logger/log_level"
	l "github.com/madeiramadeirabr/lib-go-logger/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeLogger(t *testing.T) {
	t.Run("MakeLogger", func(t *testing.T) {

		t.Run("Should return an instance of Logger", func(t *testing.T) {

			config := l.Config{
				ServiceName: "foo",
				Level:       log_level.LogLevelInfo,
			}

			logger := factory.MakeLogger(config)
			assert.IsType(t, &l.Logger{}, logger)
		})
	})
}
