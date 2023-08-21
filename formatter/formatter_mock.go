package formatter

import (
	l "github.com/madeiramadeirabr/lib-go-logger/log_level"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Format(level l.LogLevelEnum, message string, logMessageOptions LogMessageOptions) (string, error) {
	args := m.Called(level, message, logMessageOptions)

	return args.String(0), args.Error(1)
}

func (m *Mock) ClearExpectation() {
	m.ExpectedCalls = nil
}
