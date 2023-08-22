package formatter_test

import (
	"fmt"
	"github.com/madeiramadeirabr/lib-go-logger/clock"
	f "github.com/madeiramadeirabr/lib-go-logger/formatter"
	"github.com/madeiramadeirabr/lib-go-logger/log_level"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatter(t *testing.T) {
	t.Run("Formatter", func(t *testing.T) {
		formatter := f.New("lib-log-test", clock.Mock{})
		dateFixture := clock.Mock{}.GetCurrentTimestamp().Format(time.RFC3339)
		t.Run("Should format and not return error", func(t *testing.T) {
			formattedMessage, _ := formatter.Format(log_level.LogLevelError, "mensagem", f.LogMessageOptions{
				GlobalEventName: "teste",
				TraceId:         "abcdefgh",
				SessionId:       "ijklm",
			})

			responseExpected := fmt.Sprintf(
				`{"global_event_timestamp":"%s","global_event_name":"%s","level":"%s","message":"%s","service_name":"%s","trace_id":"%s","session_id":"%s"}`,
				dateFixture,
				"teste",
				log_level.LogLevelError,
				"mensagem",
				"lib-log-test",
				"abcdefgh",
				"ijklm",
			)

			assert.Equal(t, responseExpected, formattedMessage)
		})

		t.Run("Should format and omity empty", func(t *testing.T) {
			formattedMessage, _ := formatter.Format(log_level.LogLevelError, "mensagem", f.LogMessageOptions{})

			responseExpected := fmt.Sprintf(
				`{"global_event_timestamp":"%s","level":"%s","message":"%s","service_name":"%s"}`,
				dateFixture,
				log_level.LogLevelError,
				"mensagem",
				"lib-log-test",
			)

			assert.Equal(t, responseExpected, formattedMessage)
		})
	})
}
