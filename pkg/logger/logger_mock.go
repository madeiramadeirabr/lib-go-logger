package logger

import (
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) ClearExpectation() {
	m.ExpectedCalls = nil
}
