package formatter_test

import (
	"fmt"
	"lib-go-logger/pkg/clock"
	f "lib-go-logger/pkg/formatter"
	"lib-go-logger/pkg/log_level"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatter(t *testing.T) {
	t.Run("Formatter", func(t *testing.T) {
		t.Run("Should format and not return error", func(t *testing.T) {
			formatter := f.New("lib-log-test", clock.Mock{})
			dateFixture := clock.Mock{}.GetCurrentTimestamp()
			formattedMessage, _ := formatter.Format(log_level.LogLevelError, "mensagem", f.LogMessageOptions{
				GlobalEventName: "teste",
				TraceId:         "abcdefgh",
				SessionId:       "ijklm",
			})

			responseExpected := fmt.Sprintf(
				`{"global_event_timestamp":"%s","global_event_name":"%s","level":"%s","message":"%s","service_name":"%s","trace_id":"%s","session_id":"%s"}`,
				dateFixture.String(),
				"teste",
				log_level.LogLevelError,
				"mensagem",
				"lib-log-test",
				"abcdefgh",
				"ijklm",
			)

			assert.Equal(t, responseExpected, formattedMessage)
		})
	})
}
