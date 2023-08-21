package handler

import (
	l "github.com/madeiramadeirabr/lib-go-logger/log_level"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) IsHandling(level l.LogLevelEnum) bool {
	args := m.Called(level)
	return args.Bool(0)
}

func (m *Mock) Write(message string) bool {
	args := m.Called(message)
	return args.Bool(0)
}

func (m *Mock) ClearExpectation() {
	m.ExpectedCalls = nil
}
